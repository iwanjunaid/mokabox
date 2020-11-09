package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/model"

	"github.com/google/uuid"
	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventZombieAcquired(e event.EventHandler, timestamp time.Time,
	pickerGroupID uuid.UUID, originGroupID uuid.UUID,
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
