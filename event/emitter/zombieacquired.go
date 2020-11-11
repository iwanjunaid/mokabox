package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/model"

	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventZombieAcquired(e event.EventHandler, timestamp time.Time,
	pickerGroupID string, originGroupID string,
	record *model.OutboxRecord) {
	if e != nil {
		eventZombieAcquired := event.ZombieAcquired{
			PickerGroupID: pickerGroupID,
			OriginGroupID: originGroupID,
			OutboxRecord:  record,
			Timestamp:     timestamp,
		}

		e(eventZombieAcquired)
	}
}
