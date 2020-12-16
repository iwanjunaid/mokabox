# Mokabox

Mokabox is [Go](https://golang.org/) transactional outbox pattern implementation for MongoDB and Kafka.

For another implementation for Postgres and Kafka, please see [Pokabox](https://github.com/iwanjunaid/pokabox).

## Table of Contents

- [Mokabox](#mokabox)
	- [Table of Contents](#table-of-contents)
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

2. Create your app.

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/mokabox/config"
	events "github.com/iwanjunaid/mokabox/event"
	"github.com/iwanjunaid/mokabox/internal/interfaces/event"
	"github.com/iwanjunaid/mokabox/manager"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var (
		host   = "127.0.0.1"
		port   = 27017
		dbName = "basesvc"
	)

	mongoURI := fmt.Sprintf("mongodb://%s:%d/", host, port)
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	var groupID = "1f830f06-fe7c-450e-b21f-0b8569aad756"
	var bootstrapServers = "127.0.0.1:9092"
	outboxConfig := config.NewDefaultCommonOutboxConfig(groupID, dbName)

	kafkaConfig := config.NewCommonKafkaConfig(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"acks":              "all",
	})

	eventHandler := func(e event.Event) {
		switch event := e.(type) {
		case events.PickerStarted:
			fmt.Printf("%v\n", event)
		case events.Picked:
			fmt.Printf("%v\n", event)
		case events.Sent:
			fmt.Printf("%v\n", event)
		case events.StatusChanged:
			fmt.Printf("%v\n", event)
		case events.PickerPaused:
			fmt.Printf("%v\n", event)
		case events.ZombiePickerStarted:
			fmt.Printf("%v\n", event)
		case events.ZombiePicked:
			fmt.Printf("%v\n", event)
		case events.ZombieAcquired:
			fmt.Printf("%v\n", event)
		case events.ZombiePickerPaused:
			fmt.Printf("%v\n", event)
		case events.RemoverStarted:
			fmt.Printf("%v\n", event)
		case events.Removed:
			fmt.Printf("%v\n", event)
		case events.RemoverPaused:
			fmt.Printf("%v\n", event)
		case events.ErrorOccured:
			fmt.Printf("%v\n", event)
		}
	}

	manager, err := manager.New(outboxConfig, kafkaConfig, client)

	if err != nil {
		log.Fatal(err)
	}

	manager.SetEventHandler(eventHandler)
	manager.Start()
	manager.Await()
}
```
## Events Handling

TODO
