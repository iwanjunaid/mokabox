# Mokabox

Mokabox is [Go](https://golang.org/) transactional outbox pattern implementation for MongoDB and Kafka.

For another implementation for Postgres and Kafka, please see [Pokabox](https://github.com/iwanjunaid/pokabox).

## Table of Contents

- [How It Works?](#how-it-works)
- [Getting Started](#getting-started)
- [Events Handling](#events-handling)

## How It Works?

TODO

## Getting Started

1. Please make sure your outbox mongodb collection complies with mokabox's outbox model:

```go
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
```

## Events Handling

TODO
