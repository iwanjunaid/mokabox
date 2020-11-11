package manager

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/iwanjunaid/mokabox/event/emitter"
	"github.com/iwanjunaid/mokabox/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func backgroundPick(manager *CommonManager) {
	outboxConfig := manager.GetOutboxConfig()
	eventHandler := manager.GetEventHandler()
	outboxGroupID := outboxConfig.GetGroupID()

	defer manager.wg.Done()
	defer func() {
		r := recover()

		if err, ok := r.(error); ok {
			emitter.EmitEventErrorOccured(eventHandler, time.Now(), outboxGroupID, err)
			manager.wg.Add(1)
			time.Sleep(3 * time.Second)

			go backgroundPick(manager)
		}
	}()

	pollInterval := outboxConfig.GetPickerPollInterval()
	messageLimit := outboxConfig.GetPickerMessageLimitPerPoll()
	dbName := outboxConfig.GetDatabaseName()
	collectionName := outboxConfig.GetOutboxCollectionName()
	collection := manager.GetMongoClient().Database(dbName).Collection(collectionName)
	defaultContext := context.TODO()

	for {
		// Emit event PickerStarted
		emitter.EmitEventPickerStarted(eventHandler, time.Now(), outboxGroupID)

		recordsFound := false
		findOptions := options.Find()
		findOptions.SetLimit(int64(messageLimit))
		findOptions.SetSort(bson.D{
			{Key: "priority", Value: 1},
			{Key: "created_at", Value: 1},
		})

		cur, err := collection.Find(defaultContext, bson.M{
			"status": model.FlagNew,
		}, findOptions)

		if err != nil {
			panic(err)
		}

		for cur.Next(defaultContext) {
			recordsFound = true

			var record model.OutboxRecord

			recErr := cur.Decode(&record)

			if recErr != nil {
				panic(recErr)
			}

			// Emit event Picked
			emitter.EmitEventPicked(eventHandler, time.Now(), &record)

			// Send to kafka
			deliveryChan := make(chan kafka.Event, 10000)
			kafkaProducer := manager.GetKafkaProducer()

			var key []byte

			if record.KafkaKey != "" {
				key = []byte(record.KafkaKey)
			}

			errProduce := kafkaProducer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &record.KafkaTopic,
					Partition: kafka.PartitionAny},
				Value: []byte(record.KafkaValue),
				Key:   key,
			}, deliveryChan)

			if errProduce != nil {
				panic(errProduce)
			}

			kafkaEvent := <-deliveryChan
			kafkaMessage := kafkaEvent.(*kafka.Message)

			if parErr := kafkaMessage.TopicPartition.Error; parErr != nil {
				panic(parErr)
			}

			// Emit event Sent
			emitter.EmitEventSent(eventHandler, time.Now(), outboxGroupID,
				*kafkaMessage.TopicPartition.Topic,
				kafkaMessage.TopicPartition.Partition,
				&record)

			close(deliveryChan)

			update := bson.D{
				{Key: "$set", Value: bson.M{"status": "SENT"}},
			}

			result, err := collection.UpdateOne(defaultContext, bson.M{
				"_id": record.ID,
			}, update)

			if err != nil {
				panic(err)
			}

			if result.ModifiedCount > 0 {
				// Emit event StatusChanged
				emitter.EmitEventStatusChanged(eventHandler, time.Now(),
					model.FlagNew, model.FlagSent, &record)
			}
		}

		curErr := cur.Close(defaultContext)

		if curErr != nil {
			panic(curErr)
		}

		if !recordsFound {
			// Emit event PickerPaused
			emitter.EmitEventPickerPaused(eventHandler, time.Now(), outboxGroupID)

			// Pause picker
			time.Sleep(time.Duration(pollInterval) * time.Second)
		}
	}
}
