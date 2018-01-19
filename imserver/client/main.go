package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	//"log"

	"sync/atomic"

	"github.com/VolantMQ/volantmq/configuration"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// 10000 client subscribe 10000+500+500 topics
var (
	logger              *zap.Logger
	userCount           = 1
	groupCount          = 500
	publicCount         = 500
	groupTopics         []string
	publicTopics        []string
	server              = "tcp://127.0.0.1:8883"
	wg                  sync.WaitGroup
	clientCertFile      = "./client.crt"
	clientKeyFile       = "./client.key"
	testUsername        = "testuser"
	testPassword        = "testpwd"
	subPendingCount     uint64
	subSuccessCount     uint64
	receiveMessageCount uint64
	startTime           time.Time
	endTime             time.Time
)

func init() {
	logger = configuration.GetLogger().Named("client")
	for i := 1; i <= groupCount; i++ {
		groupTopics = append(groupTopics, fmt.Sprintf("g/%d", i))
	}
	for i := 1; i <= publicCount; i++ {
		publicTopics = append(publicTopics, fmt.Sprintf("p/%d", i))
	}
}

func loadTLSConfig() *tls.Config {
	return nil
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

func main() {
	tlsConfig := loadTLSConfig()
	for i := 1; i <= userCount; i++ {
		clientID := cast.ToString(i)
		go func() {
			startClient(clientID, testUsername, testPassword, tlsConfig)
		}()
	}

	watch()

	logger.Info("Spend", zap.Duration("time(ms)", endTime.Sub(startTime)/time.Millisecond))
	logger.Info("Sub Pending", zap.Uint64("count", subPendingCount))
	logger.Info("Sub Success", zap.Uint64("count", subSuccessCount))
	logger.Info("Receive message", zap.Uint64("count", receiveMessageCount))
}

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	logger.Info("Received message", zap.String("topic", message.Topic()), zap.ByteString("payload", message.Payload()))
	if receiveMessageCount == 0 {
		startTime = time.Now()
	}
	endTime = time.Now()
	atomic.AddUint64(&receiveMessageCount, 1)
}

func startClient(clientID, username, password string, tlsConfig *tls.Config) {
	done := make(chan struct{})
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
		done <- struct{}{}
	}

	client := MQTT.NewClient(connOpts)
	token := client.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		logger.Error("Connect", zap.Error(err))
	}

	<-connectC

	userTopic := "u/" + clientID
	topics := []string{userTopic}
	topics = append(topics, groupTopics...)
	topics = append(topics, publicTopics...)
	var qos byte = 1

	for _, topic := range topics {
		atomic.AddUint64(&subPendingCount, 1)
		token := client.Subscribe(topic, qos, onMessageReceived)
		token.Wait()
		if err := token.Error(); err != nil {
			logger.Error("Subscribe", zap.Error(err))
		} else {
			atomic.AddUint64(&subSuccessCount, 1)
		}
	}

	<-done
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
