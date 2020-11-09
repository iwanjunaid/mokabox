package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/model"

	"github.com/google/uuid"
	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventZombiePicked(e event.EventHandler, timestamp time.Time,
	pickerGroupID uuid.UUID, record *model.OutboxRecord) {
	if e != nil {
		eventZombiePicked := event.ZombiePicked{
			PickerGroupID: pickerGroupID,
			OutboxRecord:  record,
			Timestamp:     timestamp,
		}

		e(eventZombiePicked)
	}
}
