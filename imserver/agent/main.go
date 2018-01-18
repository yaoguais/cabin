package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

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
	logger        *zap.Logger
	grpcEndpoint  = "127.0.0.1:50051"
	httpEndPoint  = ":50050"
	redisClient   *redis.Client
	redisAddr     = "127.0.0.1:6379"
	redisPassword = ""
	redisDB       = 1
	producer      sarama.SyncProducer
	kafkaBrokers  = "localhost:9092"
	kafkaTopic    = "im"
)

type HanlderFromEndpointFunc func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

func init() {
	logger = configuration.GetLogger().Named("agent")
}

func main() {
	initRedis()
	defer closeRedis()
	initProducer()
	defer closeProducer()
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
		RegisterIMHandlerFromEndpoint,
		RegisterUserHandlerFromEndpoint,
		RegisterGroupHandlerFromEndpoint,
	}
	for _, h := range handlers {
		if err := h(ctx, mux, grpcEndpoint, opts); err != nil {
			logger.Fatal("Register handler", zap.Error(err))
		}
	}

	logger.Info("Serve http", zap.String("address", httpEndPoint))
	err := http.ListenAndServe(httpEndPoint, mux)
	if err != nil {
		logger.Fatal("Serve http", zap.Error(err))
	}
}

func startGrpcServer() {
	ln, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		logger.Fatal("Listen", zap.Error(err))
	}
	s := grpc.NewServer()

	RegisterIMServer(s, new(imServer))
	RegisterUserServer(s, new(userServer))
	RegisterGroupServer(s, new(groupServer))

	logger.Info("Serve grpc", zap.String("address", grpcEndpoint))
	err = s.Serve(ln)
	if err != nil {
		logger.Fatal("Serve grpc", zap.Error(err))
	}
}

// watch
func watch() {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGSEGV,
		syscall.SIGTERM,
		syscall.SIGSTOP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
	)

Exit:
	for {
		s := <-c
		logger.Info("Receive signal", zap.String("signal", s.String()))
		switch s {
		case syscall.SIGUSR1, syscall.SIGUSR2:
			// Do noting
		default:
			if len(c) == 0 {
				break Exit
			}
		}
	}
}

// IM
type imServer struct {
}

type jobModel struct {
	Topic    string `json:"topic"`
	Payload  string `json:"payload"`
	Qos      int32  `json:"qos"`
	Retained bool   `json:"retained"`
}

func (*imServer) Publish(ctx context.Context, req *PublishRequest) (*RetResponse, error) {
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
type userServer struct {
}

type userModel struct {
	ID  string `json:"id"`
	Pwd string `json:"pwd"`
}

func (*userServer) Add(ctx context.Context, req *AddUserRequest) (*RetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := userModel{
		ID:  req.Id,
		Pwd: req.Password,
	}
	data, err := jsoniter.Marshal(user)
	if err != nil {
		logger.Error("JSON encode", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	key := fmt.Sprintf("im:user:%s", req.Username)
	err = redisClient.Set(key, data, 0).Err()
	if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &RetResponse{
		Ok: true,
	}, nil
}

func (*userServer) Del(ctx context.Context, req *DelUserRequest) (*RetResponse, error) {
	key := fmt.Sprintf("im:user:%s", req.Username)
	err := redisClient.Del(key).Err()
	if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &RetResponse{
		Ok: true,
	}, nil
}

func (*userServer) Get(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	key := fmt.Sprintf("im:user:%s", req.Username)
	data, err := redisClient.Get(key).Bytes()
	if err == redis.Nil {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		logger.Error("Redis exec", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	var user userModel
	err = jsoniter.Unmarshal(data, &user)
	if err != nil {
		logger.Error("JSON decode", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &GetUserResponse{
		Id:       user.ID,
		Username: req.Username,
		Password: user.Pwd,
	}, nil
}

// Group
type groupServer struct {
}

func (*groupServer) AddMembers(ctx context.Context, req *AddMembersRequest) (*RetResponse, error) {
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

func (*groupServer) DelMembers(ctx context.Context, req *DelMembersRequest) (*RetResponse, error) {
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

func (*groupServer) ListMembers(ctx context.Context, req *ListMembersRequest) (*ListMembersResponse, error) {
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
