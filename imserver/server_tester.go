package main

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	//"log"

	"github.com/VolantMQ/volantmq/configuration"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// 10000 client subscribe 10000+500+500 topics
var (
	logger         *zap.Logger
	server         = "ssl://127.0.0.1:8883"
	wg             sync.WaitGroup
	clientCertFile = "./client.crt"
	clientKeyFile  = "./client.key"
	testUserID     = "2"
	testGroupID    = "2"
	testUsername   = "testuser"
	testPassword   = "testpwd"
	agentHost      = "127.0.0.1"
	agentGrpcPort  = 50051
)

func init() {
	logger = configuration.GetLogger().Named("client")
}

func main() {
	initUser(testUsername, testPassword)
	startClient(testUserID, testUsername, testPassword, loadTLSConfig())
	publishMessages()
	watch()
}

func loadTLSConfig() *tls.Config {
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		logger.Fatal("Load tls file", zap.Error(err))
	}
	return &tls.Config{
		ClientAuth:         tls.NoClientCert,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
}

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	logger.Info("Received message", zap.String("topic", message.Topic()), zap.ByteString("payload", message.Payload()))
}

func initUser(username, password string) {
	address := fmt.Sprintf("%s:%d", agentHost, agentGrpcPort)
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Grpc connect", zap.Error(err))
	}

	// Create user
	userClient := NewUserSvcClient(cc)
	_, err = userClient.Add(context.Background(), &AddUserRequest{
		Id:       testUserID,
		Username: username,
		Password: password,
		Ip:       "",
	})
	if err != nil {
		logger.Fatal("Create user", zap.Error(err))
	}

	// Join group
	groupClient := NewGroupSvcClient(cc)
	_, err = groupClient.AddMembers(context.Background(), &AddMembersRequest{
		GroupId: testGroupID,
		Members: []string{testUserID},
	})
	if err != nil {
		logger.Fatal("Join group", zap.Error(err))
	}
}

func startClient(clientID, username, password string, tlsConfig *tls.Config) {
	connectC := make(chan struct{})
	connOpts := MQTT.
		NewClientOptions().
		AddBroker(server).
		SetClientID(clientID).
		SetCleanSession(true).
		SetUsername(username).
		SetPassword(password)

	if tlsConfig != nil {
		connOpts.SetTLSConfig(tlsConfig)
	}

	connOpts.OnConnect = func(c MQTT.Client) {
		logger.Info("Connect", zap.String("clientID", clientID))
		connectC <- struct{}{}
	}
	connOpts.OnConnectionLost = func(client MQTT.Client, err error) {
		logger.Info("Connection lost", zap.Error(err))
	}

	client := MQTT.NewClient(connOpts)
	token := client.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		logger.Error("Connect", zap.Error(err))
	}

	logger.Info("Connect...")
	<-connectC

	topics := []string{
		"u/" + clientID,
		"g/" + testGroupID,
		"p/" + clientID,
	}
	var qos byte = 1

	for _, topic := range topics {
		token := client.Subscribe(topic, qos, onMessageReceived)
		token.Wait()
		if err := token.Error(); err != nil {
			logger.Error("Subscribe", zap.Error(err))
		}
	}
}

func publishMessages() {
	address := fmt.Sprintf("%s:%d", agentHost, agentGrpcPort)
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Grpc connect", zap.Error(err))
	}
	// Publish user message
	imClient := NewIMSvcClient(cc)
	_, err = imClient.Publish(context.Background(), &PublishRequest{
		Topic:    "u/" + testUserID,
		Payload:  time.Now().Format("2006-01-02 15:04:05.999999999"),
		Qos:      1,
		Retained: false,
	})
	if err != nil {
		logger.Fatal("Publish user message", zap.Error(err))
	}

	// Publish group message
	_, err = imClient.Publish(context.Background(), &PublishRequest{
		Topic:    "g/" + testGroupID,
		Payload:  time.Now().Format("2006-01-02 15:04:05.999999999"),
		Qos:      1,
		Retained: false,
	})
	if err != nil {
		logger.Fatal("Publish group message", zap.Error(err))
	}

	// Publish public message
	_, err = imClient.Publish(context.Background(), &PublishRequest{
		Topic:    "p/" + testUserID,
		Payload:  time.Now().Format("2006-01-02 15:04:05.999999999"),
		Qos:      1,
		Retained: false,
	})
	if err != nil {
		logger.Fatal("Publish public message", zap.Error(err))
	}
}
