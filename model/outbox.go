package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	FlagNew  = "NEW"
	FlagSent = "SENT"
)

type OutboxRecord struct {
	ID         primitive.ObjectID `bson:"_id"`
	GroupID    string             `bson:"group_id"`
	KafkaTopic string             `bson:"kafka_topic"`
	KafkaKey   string             `bson:"kafka_key"`
	KafkaValue string             `bson:"kafka_value"`
	Priority   uint               `bson:"priority"`
	Status     string             `bson:"status"`
	Version    uint               `bson:"version"`
	CreatedAt  time.Time          `bson:"created_at"`
	SentAt     time.Time          `bson:"sent_at"`
}
