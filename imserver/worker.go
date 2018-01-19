package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/VolantMQ/volantmq/configuration"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	cfgFile          string
	logger           *zap.Logger
	workerjobSize    int
	workerNumber     int
	internalClientID string
	internalUsername string
	internalPassword string
	agentHost        string
	agentGrpcPort    int
	kafkaBrokers     string
	kafkaTopic       string
)

type BrokerManager struct {
	brokerClient  BrokerSvcClient
	brokerWorkers []*BrokerWorker
	done          chan struct{}
	sync.Mutex
}

type BrokerWorker struct {
	broker       Broker
	workers      []*Worker
	consumerID   string
	consumer     sarama.Consumer
	consumerDone chan struct{}
	done         chan struct{}
	shutdown     chan struct{}
	wg           sync.WaitGroup
}

type Worker struct {
	client  MQTT.Client
	jobChan chan jobModel
}

func init() {
	logger = configuration.GetLogger().Named("worker")

	flag.StringVar(&cfgFile, "c", "", "config file")
	flag.Parse()

	cfgData, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		panic(err)
	}

	workerjobSize = jsoniter.Get(cfgData, "worker", "jobSize").ToInt()
	workerNumber = jsoniter.Get(cfgData, "worker", "number").ToInt()
	internalClientID = jsoniter.Get(cfgData, "worker", "id").ToString()
	internalUsername = jsoniter.Get(cfgData, "worker", "username").ToString()
	internalPassword = jsoniter.Get(cfgData, "worker", "password").ToString()
	agentHost = jsoniter.Get(cfgData, "agent", "host").ToString()
	agentGrpcPort = jsoniter.Get(cfgData, "agent", "ports", "grpc").ToInt()
	kafkaBrokers = jsoniter.Get(cfgData, "kafka", "brokers").ToString()
	kafkaTopic = jsoniter.Get(cfgData, "kafka", "topic").ToString()

	if workerjobSize <= 0 {
		workerjobSize = 16
	}
	if workerNumber <= 0 {
		workerNumber = 16
	}
}

func main() {
	bm := NewBrokerManager()
	if err := bm.start(); err != nil {
		panic(err)
	}
	watch()
	bm.close()
}

// BrokerManager

func NewBrokerManager() *BrokerManager {
	return &BrokerManager{
		done: make(chan struct{}),
	}
}

func (bm *BrokerManager) start() error {
	logger.Info("Broker manager start")
	address := fmt.Sprintf("%s:%d", agentHost, agentGrpcPort)
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	bm.brokerClient = NewBrokerSvcClient(cc)

	go func() {
		bm.fetchBrokers()
		t := time.NewTicker(30 * time.Second)

	Exit:
		for {
			select {
			case <-t.C:
				bm.fetchBrokers()
			case <-bm.done:
				break Exit
			}
		}

		for _, v := range bm.brokerWorkers {
			v.close()
		}
	}()

	return nil
}

func (bm *BrokerManager) close() {
	close(bm.done)
	for _, v := range bm.brokerWorkers {
		<-v.shutdown
	}
	logger.Info("Broker manager done")
}

func (bm *BrokerManager) fetchBrokers() {
	resp, err := bm.brokerClient.List(context.Background(), &ListBrokerRequest{})
	if err != nil {
		logger.Error("List Brokers", zap.Error(err))
		return
	}

	bm.Lock()
	defer bm.Unlock()

	var nowBrokerWorkers []*BrokerWorker
	brokerMap := make(map[string]struct{})
	for _, v := range resp.Brokers {
		brokerMap[v.External] = struct{}{}
	}

	prevBrokerMap := make(map[string]*BrokerWorker)
	for _, v := range bm.brokerWorkers {
		prevBrokerMap[v.broker.External] = v
		_, ok := brokerMap[v.broker.External]
		if ok {
			nowBrokerWorkers = append(nowBrokerWorkers, v)
		} else {
			v.close()
		}
	}

	for _, v := range resp.Brokers {
		_, ok := prevBrokerMap[v.External]
		if !ok {
			nbw := NewBrokerWorker(*v)
			err := nbw.start()
			if err != nil {
				logger.Error("Broker Worker", zap.Error(err), zap.String("broker", v.Internal))
			} else {
				nowBrokerWorkers = append(nowBrokerWorkers, nbw)
			}
		}
	}

	bm.brokerWorkers = nowBrokerWorkers
}

// BrokerWorker

func NewBrokerWorker(broker Broker) *BrokerWorker {
	consumeID := strings.Replace(broker.External, ".", "_", -1)
	consumeID = strings.Replace(consumeID, ":", "_", -1)
	consumeID = strings.Replace(consumeID, "/", "_", -1)

	return &BrokerWorker{
		broker:       broker,
		consumerID:   consumeID,
		consumerDone: make(chan struct{}),
		shutdown:     make(chan struct{}),
		done:         make(chan struct{}),
	}
}

func (bw *BrokerWorker) start() error {
	// Create workers
	workers, err := bw.createWorkers()
	if err != nil {
		return err
	}
	bw.workers = workers

	logger.Info("Broker worker start")

	// Start Workers
	for _, v := range bw.workers {
		worker := v
		bw.wg.Add(1)
		go func() {
			defer bw.wg.Done()
			bw.startWorker(worker)
		}()
	}

	// Create consume
	consumer, err := bw.newConsumer()
	if err != nil {
		return err
	}
	bw.consumer = consumer

	// Consume job
	bw.wg.Add(1)
	go func() {
		defer bw.wg.Done()
		bw.consume()
	}()

	go func() {
		bw.wg.Wait()
		close(bw.shutdown)
		logger.Info("Broker worker done")
	}()

	return nil
}

func (bw *BrokerWorker) createWorkers() ([]*Worker, error) {
	var workers []*Worker
	for i := 0; i < workerNumber; i++ {
		clientID := fmt.Sprintf("%s-%s-%d", internalClientID, bw.consumerID, i)
		broker := fmt.Sprintf("tcp://%s", bw.broker.Internal)
		client, err := createMQTTClient(broker, clientID, internalUsername, internalPassword)
		if err != nil {
			for _, v := range workers {
				v.close()
			}
			return nil, err
		}
		workers = append(workers, NewWorker(client))
	}
	return workers, nil
}

func (bw *BrokerWorker) startWorker(worker *Worker) {
Exit:
	for {
		select {
		case job := <-worker.jobChan:
			token := worker.client.Publish(job.Topic, byte(job.Qos), job.Retained, job.Payload)
			token.Wait()
			if err := token.Error(); err != nil {
				data, _ := jsoniter.Marshal(job)
				logger.Error("Publish", zap.Error(err), zap.ByteString("job", data))
			}
		case <-bw.done:
			break Exit
		}
	}

	for {
		if _, ok := <-bw.consumerDone; !ok && len(worker.jobChan) == 0 {
			break
		}

		select {
		case job := <-worker.jobChan:
			token := worker.client.Publish(job.Topic, byte(job.Qos), job.Retained, job.Payload)
			token.Wait()
			if err := token.Error(); err != nil {
				data, _ := jsoniter.Marshal(job)
				logger.Error("Publish rest", zap.Error(err), zap.ByteString("job", data))
			}
		default:
		}
	}
}

func (bw *BrokerWorker) newConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.ClientID = bw.consumerID
	config.ChannelBufferSize = workerjobSize

	brokerList := strings.Split(kafkaBrokers, ",")
	return sarama.NewConsumer(brokerList, config)
}

func (bw *BrokerWorker) consume() {
	partitionList, err := bw.consumer.Partitions(kafkaTopic)
	if err != nil {
		logger.Error("List kafka partition", zap.Error(err))
		return
	}

	var wg sync.WaitGroup
	for _, v := range partitionList {
		partition := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			bw.consumePartition(partition)
		}()
	}
	wg.Wait()

	close(bw.consumerDone)
}

func (bw *BrokerWorker) consumePartition(partition int32) {
	pc, err := bw.consumer.ConsumePartition(kafkaTopic, partition, sarama.OffsetOldest)
	if err != nil {
		logger.Error("Kafka partition consumer", zap.Error(err))
		return
	}

	errs := pc.Errors()
	bw.wg.Add(1)
	go func() {
		defer bw.wg.Done()
	Exit:
		for {
			select {
			case err := <-errs:
				if err != nil {
					logger.Error("Protition consumer", zap.Error(err))
				}
			case <-bw.done:
				break Exit
			}
		}
	}()

	msgs := pc.Messages()
Exit:
	for {
		select {
		case msg := <-msgs:
			bw.deliverJobToWorker(msg.Value)
		case <-bw.done:
			break Exit
		}
	}

	pc.AsyncClose()

	for msg := range msgs {
		bw.deliverJobToWorker(msg.Value)
	}
}

func (bw *BrokerWorker) deliverJobToWorker(data []byte) {
	workerCount := len(bw.workers)
	if workerCount == 0 {
		logger.Error("No worker")
		return
	}

	var job jobModel
	err := jsoniter.Unmarshal(data, &job)
	if err != nil {
		logger.Error("JSON decode", zap.ByteString("data", data), zap.Error(err))
		return
	}

	ts := strings.Split(job.Topic, "/")
	foundNumber := false
	var idx int

	for _, v := range ts {
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			if i < 0 {
				i = -i
			}
			idx = int(i) % workerCount
			foundNumber = true
			break
		}
	}

	if !foundNumber {
		idx = rand.Intn(workerCount)
	}

	bw.workers[idx].jobChan <- job
}

func (bw *BrokerWorker) close() {
	close(bw.done)
}

// Worker

func NewWorker(client MQTT.Client) *Worker {
	return &Worker{
		client:  client,
		jobChan: make(chan jobModel, workerjobSize),
	}
}

func (w *Worker) close() {
	close(w.jobChan)
	w.client.Disconnect(100)
}

// MQTT Client

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
	if !client.IsConnected() {
		return nil, errors.New("Not connected")
	}

	return client, nil
}
