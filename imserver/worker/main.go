package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"

	"math/rand"

	"github.com/Shopify/sarama"
	"github.com/VolantMQ/volantmq/configuration"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/json-iterator/go"
	"go.uber.org/zap"
)

var (
	logger           *zap.Logger
	workerNumber     = 16
	kafkaConsumer    sarama.Consumer
	kafkaBrokers     = "localhost:9092"
	kafkaTopic       = "im"
	jobChan          = make(chan []byte, 10000)
	mqttServer       = "tcp://127.0.0.1:1883"
	mqttClientID     = "robot"
	internalUsername = "testusername"
	internalPassword = "testpassword"
	mqttWorkers      []*mqttWorker
	mqttClientNumber = 4
	jobBufferSize    = mqttClientNumber
	doneChan         = make(chan struct{})
	doneOKChan       = make(chan struct{})
	receiveJobCount  uint64
	publishJobCount  uint64
	jobLogInterval   uint64 = 1000
)

type jobModel struct {
	Topic    string `json:"topic"`
	Payload  string `json:"payload"`
	Qos      int32  `json:"qos"`
	Retained bool   `json:"retained"`
}

type mqttWorker struct {
	id      int
	client  MQTT.Client
	jobChan chan jobModel
}

func init() {
	logger = configuration.GetLogger().Named("worker")
}

func main() {
	initMQTTWorker()
	defer closeMQTTWorker()
	go produce()

	go consumeJob()

	initConsumer()
	defer closeConsumer()
	go consume()

	watch()
	close(doneChan)
	<-doneOKChan
	logger.Info("Server done")
}

func initMQTTWorker() {
	for i := 0; i < mqttClientNumber; i++ {
		clientID := fmt.Sprintf("%s-%d", mqttClientID, i)
		client, err := createMQTTClient(mqttServer, clientID, internalUsername, internalPassword)
		if err != nil {
			logger.Fatal("Create MQTT client", zap.Error(err))
		}
		mw := &mqttWorker{
			id:      i,
			client:  client,
			jobChan: make(chan jobModel, jobBufferSize),
		}
		mqttWorkers = append(mqttWorkers, mw)
	}
}

func closeMQTTWorker() {
	for _, mw := range mqttWorkers {
		mw.client.Disconnect(100)
	}
}

func consumeJob() {
	for v := range jobChan {
		var job jobModel
		err := jsoniter.Unmarshal(v, &job)
		if err != nil {
			logger.Error("JSON decode", zap.Error(err))
			continue
		}
		ts := strings.Split(job.Topic, "/")
		idx := rand.Intn(workerNumber)
		if len(ts) == 2 {
			if i, err := strconv.ParseInt(ts[1], 10, 64); err == nil {
				if i < 0 {
					i = -i
				}
				idx = int(i) % workerNumber
			}
		}
		mw := mqttWorkers[idx]
		mw.jobChan <- job
	}
}

func produce() {
	var wg sync.WaitGroup
	for _, v := range mqttWorkers {
		wg.Add(1)
		mw := v
		go func() {
			defer wg.Done()
			mqttWorkerProduce(mw)
		}()
	}
	wg.Wait()
	logger.Info("Publish job", zap.Uint64("count", publishJobCount))
	doneOKChan <- struct{}{}
}

func mqttWorkerProduce(mw *mqttWorker) {
	logger.Info("MQTT worker start", zap.Int("id", mw.id))

Exit:
	for {
		select {
		case job := <-mw.jobChan:
			token := mw.client.Publish(job.Topic, byte(job.Qos), job.Retained, job.Payload)
			token.Wait()
			if err := token.Error(); err != nil {
				data, _ := jsoniter.Marshal(job)
				logger.Error("Publish", zap.Error(err), zap.ByteString("job", data))
			}

			atomic.AddUint64(&publishJobCount, 1)
			if publishJobCount%jobLogInterval == 0 {
				logger.Info("Publish current", zap.Uint64("count", publishJobCount))
			}
		case <-doneChan:
			break Exit
		}
	}

	for {
		_, ok := <-jobChan
		if !ok && len(mw.jobChan) == 0 {
			break // Consume all the jobs
		}

		select {
		case job := <-mw.jobChan:
			token := mw.client.Publish(job.Topic, byte(job.Qos), job.Retained, job.Payload)
			token.Wait()
			if err := token.Error(); err != nil {
				data, _ := jsoniter.Marshal(job)
				logger.Error("Publish rest", zap.Error(err), zap.ByteString("job", data))
			}

			atomic.AddUint64(&publishJobCount, 1)
			if publishJobCount%jobLogInterval == 0 {
				logger.Info("Publish current", zap.Uint64("count", publishJobCount))
			}
		default:
		}
	}

	logger.Info("MQTT worker finish", zap.Int("id", mw.id))
}

func createMQTTClient(broker, clientID, username, password string) (MQTT.Client, error) {
	connOpts := MQTT.
		NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID).
		SetCleanSession(true).
		SetUsername(username).
		SetPassword(password)

	connOpts.OnConnect = func(c MQTT.Client) {
		logger.Info("Connect", zap.String("clientID", clientID))
	}
	connOpts.OnConnectionLost = func(client MQTT.Client, err error) {
		logger.Info("Connection lost", zap.String("clientID", clientID), zap.Error(err))
	}

	client := MQTT.NewClient(connOpts)
	token := client.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		logger.Error("Connect", zap.Error(err))
		return nil, err
	}

	return client, nil
}

func initConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.ClientID = "worker"
	config.ChannelBufferSize = jobBufferSize

	var err error
	brokerList := strings.Split(kafkaBrokers, ",")
	kafkaConsumer, err = sarama.NewConsumer(brokerList, config)
	if err != nil {
		logger.Fatal("Kafka consumer", zap.Error(err))
	}
}

func closeConsumer() {
	err := kafkaConsumer.Close()
	if err != nil {
		logger.Error("Close consumer", zap.Error(err))
	}
}

func consume() {
	partitionList, err := kafkaConsumer.Partitions(kafkaTopic)
	if err != nil {
		logger.Fatal("Kafka partitions", zap.Error(err))
	}

	var wg sync.WaitGroup
	for i, v := range partitionList {
		wg.Add(1)
		id := i
		partition := v
		go func() {
			defer wg.Done()
			logger.Info("Partition consumer start", zap.Int("id", id), zap.Int32("partition", partition))
			consumePartition(partition)
			logger.Info("Partition consumer finish", zap.Int("id", id), zap.Int32("partition", partition))
		}()
	}
	wg.Wait()
	logger.Info("Receive job", zap.Uint64("count", receiveJobCount))

	close(jobChan)
}

func consumePartition(partition int32) error {
	pc, err := kafkaConsumer.ConsumePartition(kafkaTopic, partition, sarama.OffsetOldest)
	if err != nil {
		logger.Error("Kafka partition consumer", zap.Error(err))
		return err
	}

	errs := pc.Errors()
	go func() {
	Exit:
		for {
			select {
			case err := <-errs:
				if err != nil {
					logger.Error("Protition consumer", zap.Error(err))
				}
			case <-doneChan:
				break Exit
			}
		}
	}()

	msgs := pc.Messages()
Exit:
	for {
		select {
		case msg := <-msgs:
			jobChan <- msg.Value
			atomic.AddUint64(&receiveJobCount, 1)
			if receiveJobCount%jobLogInterval == 0 {
				logger.Info("Receive current", zap.Uint64("count", receiveJobCount))
			}
		case <-doneChan:
			break Exit
		}
	}

	pc.AsyncClose()

	for msg := range msgs {
		jobChan <- msg.Value
		atomic.AddUint64(&receiveJobCount, 1)
		if receiveJobCount%jobLogInterval == 0 {
			logger.Info("Receive current", zap.Uint64("count", receiveJobCount))
		}
	}

	return nil
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
