package manager

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/mokabox/event"
	"github.com/iwanjunaid/mokabox/pkg/interfaces/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Manager interface {
	GetOutboxConfig() config.OutboxConfig
	GetKafkaConfig() config.KafkaConfig
	GetKafkaProducer() *kafka.Producer
	GetMongoClient() *mongo.Client
	SetEventHandler(event.EventHandler)
	GetEventHandler() event.EventHandler

	// TODO: Create insert function to persist to mongodb
	// Insert(*sql.Tx, *model.OutboxRecord) error
	Start() error
	Await()
}
