package manager

import (
	"context"
	"time"

	"github.com/iwanjunaid/mokabox/event/emitter"
	"github.com/iwanjunaid/mokabox/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func backgroundZombiePick(manager *CommonManager) {
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

	pollInterval := outboxConfig.GetZombiePickerPollInterval()
	messageLimit := outboxConfig.GetZombiePickerMessageLimitPerPoll()
	zombieInterval := outboxConfig.GetZombieInterval()
	dbName := outboxConfig.GetDatabaseName()
	collectionName := outboxConfig.GetOutboxCollectionName()
	collection := manager.GetMongoClient().Database(dbName).Collection(collectionName)
	defaultContext := context.TODO()

	for {
		// Emit event ZombiePickerStarted
		emitter.EmitEventZombiePickerStarted(eventHandler, time.Now(), outboxGroupID)

		recordsFound := false
		findOptions := options.Find()
		findOptions.SetLimit(int64(messageLimit))
		timeAgo := time.Now().Add(time.Duration(-zombieInterval) * time.Second)

		cur, err := collection.Find(defaultContext, bson.M{
			"group_id": bson.M{
				"$ne": outboxGroupID,
			},
			"created_at": bson.M{
				"$lt": timeAgo,
			},
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

			// Emit event ZombiePicked
			emitter.EmitEventZombiePicked(eventHandler, time.Now(),
				outboxGroupID, &record)

			update := bson.D{
				{Key: "$set", Value: bson.M{
					"group_id": outboxGroupID,
					"sent_at":  time.Now(),
				}},
				{Key: "$inc", Value: bson.M{"version": 1}},
			}

			result, err := collection.UpdateOne(defaultContext, bson.M{
				"_id":     record.ID,
				"version": record.Version,
			}, update)

			if err != nil {
				panic(err)
			}

			if result.ModifiedCount > 0 {
				// Emit event ZombieAcquired
				emitter.EmitEventZombieAcquired(eventHandler, time.Now(),
					outboxGroupID, record.GroupID, &record)
			}
		}

		curErr := cur.Close(defaultContext)

		if curErr != nil {
			panic(curErr)
		}

		if !recordsFound {
			// Emit event ZombiePickerPaused
			emitter.EmitEventZombiePickerPaused(eventHandler, time.Now(), outboxGroupID)

			// Pause zombie picker
			time.Sleep(time.Duration(pollInterval) * time.Second)
		}
	}
}
