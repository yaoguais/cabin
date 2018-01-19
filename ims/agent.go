package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/VolantMQ/volantmq/configuration"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/json-iterator/go"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/redis.v5"
)

var (
	cfgFile       string
	logger        *zap.Logger
	redisClient   *redis.Client
	grpcPort      int
	httpPort      int
	redisAddr     string
	redisPassword string
	redisDB       int
	producer      sarama.SyncProducer
	kafkaBrokers  string
	kafkaTopic    string
	brokersCache  []*Broker
)

type HanlderFromEndpointFunc func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

func init() {
	logger = configuration.GetLogger().Named("agent")

	flag.StringVar(&cfgFile, "c", "", "config file")
	flag.Parse()

	cfgData, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		panic(err)
	}

	grpcPort = jsoniter.Get(cfgData, "agent", "ports", "grpc").ToInt()
	httpPort = jsoniter.Get(cfgData, "agent", "ports", "http").ToInt()
	redisAddr = jsoniter.Get(cfgData, "agent", "redis", "address").ToString()
	redisPassword = jsoniter.Get(cfgData, "agent", "redis", "password").ToString()
	redisDB = jsoniter.Get(cfgData, "agent", "redis", "db").ToInt()
	kafkaBrokers = jsoniter.Get(cfgData, "kafka", "brokers").ToString()
	kafkaTopic = jsoniter.Get(cfgData, "kafka", "topic").ToString()
}

func main() {
	initRedis()
	defer closeRedis()
	initProducer()
	defer closeProducer()
	go updateBrokersCache()
	go startGrpcServer()
	go startHttpServer()
	watch()
}

func initRedis() {
	options := &redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	}
	redisClient = redis.NewClient(options)
	err := redisClient.Ping().Err()
	if err != nil {
		logger.Fatal("Redis connect", zap.Error(err))
	}
}

func closeRedis() {
	err := redisClient.Close()
	if err != nil {
		logger.Error("Close redis", zap.Error(err))
	}
}

func initProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	var err error
	brokerList := strings.Split(kafkaBrokers, ",")
	producer, err = sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		logger.Fatal("Kafka producer", zap.Error(err))
	}
}

func closeProducer() {
	err := producer.Close()
	if err != nil {
		logger.Error("Close producer", zap.Error(err))
	}
}

func startHttpServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	handlers := []HanlderFromEndpointFunc{
		RegisterBrokerSvcHandlerFromEndpoint,
		RegisterIMSvcHandlerFromEndpoint,
		RegisterUserSvcHandlerFromEndpoint,
		RegisterGroupSvcHandlerFromEndpoint,
	}
	grpcEndpoint := fmt.Sprintf("127.0.0.1:%d", grpcPort)
	for _, h := range handlers {
		if err := h(ctx, mux, grpcEndpoint, opts); err != nil {
			logger.Fatal("Register handler", zap.Error(err))
		}
	}

	endpoint := fmt.Sprintf(":%d", httpPort)
	logger.Info("Serve http", zap.String("address", endpoint))
	err := http.ListenAndServe(endpoint, mux)
	if err != nil {
		logger.Fatal("Serve http", zap.Error(err))
	}
}

func startGrpcServer() {
	endpoint := fmt.Sprintf(":%d", grpcPort)
	ln, err := net.Listen("tcp", endpoint)
	if err != nil {
		logger.Fatal("Listen", zap.Error(err))
	}
	s := grpc.NewServer()

	RegisterBrokerSvcServer(s, new(brokerSvc))
	RegisterIMSvcServer(s, new(imSvc))
	RegisterUserSvcServer(s, new(userSvc))
	RegisterGroupSvcServer(s, new(groupSvc))

	logger.Info("Serve grpc", zap.String("address", endpoint))
	err = s.Serve(ln)
	if err != nil {
		logger.Fatal("Serve grpc", zap.Error(err))
	}
}

// Broker
type brokerSvc struct {
}

func saveBroker(broker Broker) error {
	data, err := jsoniter.Marshal(broker)
	if err != nil {
		logger.Error("JSON encode", zap.Error(err))
		return err
	}

	hkey := "im:brokers"
	err = redisClient.HSet(hkey, broker.External, data).Err()
	if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return err
	}
	return nil
}

func listBroker() ([]*Broker, error) {
	var brokers []*Broker
	hkey := "im:brokers"
	vals, err := redisClient.HVals(hkey).Result()
	if err != nil {
		return nil, err
	}
	// logger.Info("List brokers", zap.Strings("brokers", vals))
	for _, v := range vals {
		var broker Broker
		err = jsoniter.Unmarshal([]byte(v), &broker)
		if err != nil {
			logger.Error("JSON decode", zap.Error(err))
			return nil, err
		}
		brokers = append(brokers, &broker)
	}
	return brokers, nil
}

func brokerForUser(user User) *Broker {
	now := time.Now().Unix()
	for _, v := range brokersCache {
		if now-v.UpdateTime < 300 && (v.MaxConn < 0 || int64(v.Conn) < v.MaxConn-100) {
			return v
		}
	}
	return nil
}

func updateBrokersCache() {
	brokers, err := listBroker()
	if err != nil {
		logger.Fatal("List brokers", zap.Error(err))
	}
	brokersCache = brokers

	for range time.NewTicker(time.Second).C {
		brokers, err := listBroker()
		if err != nil {
			logger.Error("List brokers", zap.Error(err))
		} else {
			brokersCache = brokers
		}
	}
}

func (*brokerSvc) Update(ctx context.Context, req *UpdateBrokerRequest) (*Broker, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	broker := Broker{
		External:   req.External,
		Internal:   req.Internal,
		Conn:       req.Conn,
		MaxConn:    req.MaxConn,
		UpdateTime: time.Now().Unix(),
	}

	err := saveBroker(broker)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &broker, nil
}

func (*brokerSvc) List(context.Context, *ListBrokerRequest) (*ListBrokerResponse, error) {
	brokers, err := listBroker()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ListBrokerResponse{
		Brokers: brokers,
	}, nil
}

// IM
type imSvc struct {
}

func (*imSvc) Publish(ctx context.Context, req *PublishRequest) (*RetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	job := jobModel{
		Topic:    req.Topic,
		Payload:  req.Payload,
		Qos:      req.Qos,
		Retained: req.Retained,
	}
	data, err := jsoniter.Marshal(job)
	if err != nil {
		logger.Error("JSON encode", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	msg := &sarama.ProducerMessage{
		Topic: kafkaTopic,
		Value: sarama.ByteEncoder(data),
	}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		logger.Error("Kafka send message", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &RetResponse{
		Ok: true,
	}, nil
}

// User
type userSvc struct {
}

func saveUser(user User) error {
	data, err := jsoniter.Marshal(user)
	if err != nil {
		logger.Error("JSON encode", zap.Error(err))
		return err
	}

	key := fmt.Sprintf("im:user:%s", user.Username)
	err = redisClient.Set(key, data, 0).Err()
	if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return err
	}
	return nil
}

func getUserByUsername(username string) (*User, error) {
	key := fmt.Sprintf("im:user:%s", username)
	data, err := redisClient.Get(key).Bytes()
	if err == redis.Nil {
		return nil, err
	} else if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return nil, err
	}

	var user User
	err = jsoniter.Unmarshal(data, &user)
	if err != nil {
		logger.Error("JSON decode", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func delUserByUsername(username string) error {
	key := fmt.Sprintf("im:user:%s", username)
	err := redisClient.Del(key).Err()
	if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return err
	}
	return nil
}

func (*userSvc) Add(ctx context.Context, req *AddUserRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := User{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Ip:       req.Ip,
		Broker:   "",
	}
	if broker := brokerForUser(user); broker != nil {
		user.Broker = broker.External
	}

	err := saveUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user, nil
}

func (*userSvc) Del(ctx context.Context, req *DelUserRequest) (*User, error) {
	user, err := getUserByUsername(req.Username)
	if err != nil {
		if err == redis.Nil {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = delUserByUsername(req.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return user, nil
}

func (*userSvc) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	user, err := getUserByUsername(req.Username)
	if err != nil {
		if err == redis.Nil {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return user, nil
}

// Group
type groupSvc struct {
}

func (*groupSvc) AddMembers(ctx context.Context, req *AddMembersRequest) (*RetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	key := fmt.Sprintf("im:group:%s", req.GroupId)
	err := redisClient.SAdd(key, sliceStringToInterface(req.Members)...).Err()
	if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &RetResponse{
		Ok: true,
	}, nil
}

func sliceStringToInterface(s []string) []interface{} {
	t := make([]interface{}, len(s))
	for i, v := range s {
		t[i] = v
	}
	return t
}

func (*groupSvc) DelMembers(ctx context.Context, req *DelMembersRequest) (*RetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	key := fmt.Sprintf("im:group:%s", req.GroupId)
	err := redisClient.SRem(key, sliceStringToInterface(req.Members)...).Err()
	if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &RetResponse{
		Ok: true,
	}, nil
}

func (*groupSvc) IsMember(ctx context.Context, req *IsMemberRequest) (*RetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	key := fmt.Sprintf("im:group:%s", req.GroupId)
	ok, err := redisClient.SIsMember(key, req.MemberId).Result()
	if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &RetResponse{
		Ok: ok,
	}, nil
}

func (*groupSvc) ListMembers(ctx context.Context, req *ListMembersRequest) (*ListMembersResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	key := fmt.Sprintf("im:group:%s", req.GroupId)
	members, err := redisClient.SMembers(key).Result()
	if err == redis.Nil {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ListMembersResponse{
		Members: members,
	}, nil
}
