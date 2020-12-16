package manager

import (
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/mokabox/event"
	"github.com/iwanjunaid/mokabox/internal/interfaces/config"
	"github.com/iwanjunaid/mokabox/internal/interfaces/manager"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommonManager struct {
	outboxConfig  config.OutboxConfig
	kafkaConfig   config.KafkaConfig
	kafkaProducer *kafka.Producer
	eventHandler  event.EventHandler
	client        *mongo.Client
	wg            *sync.WaitGroup
}

func New(outboxConfig config.OutboxConfig, kafkaConfig config.KafkaConfig, client *mongo.Client) (manager.Manager, error) {
	var wg sync.WaitGroup

	kafkaProducer, err := kafka.NewProducer(kafkaConfig.GetConfigMap())

	if err != nil {
		return nil, err
	}

	manager := &CommonManager{
		outboxConfig:  outboxConfig,
		kafkaConfig:   kafkaConfig,
		kafkaProducer: kafkaProducer,
		eventHandler:  nil,
		client:        client,
		wg:            &wg,
	}

	return manager, nil
}

func (m *CommonManager) GetOutboxConfig() config.OutboxConfig {
	return m.outboxConfig
}

func (m *CommonManager) GetKafkaConfig() config.KafkaConfig {
	return m.kafkaConfig
}

func (m *CommonManager) GetKafkaProducer() *kafka.Producer {
	return m.kafkaProducer
}

func (m *CommonManager) SetEventHandler(e event.EventHandler) {
	m.eventHandler = e
}

func (m *CommonManager) GetEventHandler() event.EventHandler {
	return m.eventHandler
}

func (m *CommonManager) GetMongoClient() *mongo.Client {
	return m.client
}

func (m *CommonManager) Start() error {
	m.wg.Add(2)

	go backgroundPick(m)
	go backgroundZombiePick(m)
	// go backgroundRemove(m)

	return nil
}

func (m *CommonManager) Await() {
	m.wg.Wait()
}
