package manager

import (
	"context"
	"time"

	"github.com/iwanjunaid/mokabox/event/emitter"
	"github.com/iwanjunaid/mokabox/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func backgroundRemove(manager *CommonManager) {
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

			go backgroundRemove(manager)
		}
	}()

	pollInterval := outboxConfig.GetRemoverPollInterval()
	messageLimit := outboxConfig.GetRemoverMessageLimitPerPoll()
	dbName := outboxConfig.GetDatabaseName()
	collectionName := outboxConfig.GetOutboxCollectionName()
	collection := manager.GetMongoClient().Database(dbName).Collection(collectionName)
	defaultContext := context.TODO()

	for {
		// Emit event RemoverStarted
		emitter.EmitEventRemoverStarted(eventHandler, time.Now(), outboxGroupID)

		recordsFound := false
		findOptions := options.Find()
		findOptions.SetLimit(int64(messageLimit))

		cur, err := collection.Find(defaultContext, bson.M{
			"status": model.FlagSent,
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

			filter := bson.D{primitive.E{Key: "_id", Value: record.ID}}
			result, err := collection.DeleteOne(defaultContext, filter)

			if err != nil {
				panic(err)
			}

			if result.DeletedCount > 0 {
				// Emit event Removed
				emitter.EmitEventRemoved(eventHandler, time.Now(), outboxGroupID, &record)
			}
		}

		curErr := cur.Close(defaultContext)

		if curErr != nil {
			panic(curErr)
		}

		if !recordsFound {
			// Emit event RemoverPaused
			emitter.EmitEventRemoverPaused(eventHandler, time.Now(), outboxGroupID)

			// Pause remover
			time.Sleep(time.Duration(pollInterval) * time.Second)
		}
	}
}
