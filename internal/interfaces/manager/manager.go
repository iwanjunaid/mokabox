package manager

import (
	"database/sql"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/mokabox/event"
	"github.com/iwanjunaid/mokabox/internal/interfaces/config"
)

type Manager interface {
	GetOutboxConfig() config.OutboxConfig
	GetKafkaConfig() config.KafkaConfig
	GetKafkaProducer() *kafka.Producer
	GetDB() *sql.DB
	SetEventHandler(event.EventHandler)
	GetEventHandler() event.EventHandler
	// Insert(*sql.Tx, *model.OutboxRecord) error
	Start() error
	Await()
}
