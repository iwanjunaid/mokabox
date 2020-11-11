package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
	"github.com/iwanjunaid/mokabox/model"
)

func EmitEventZombiePicked(e event.EventHandler, timestamp time.Time,
	pickerGroupID string, record *model.OutboxRecord) {
	if e != nil {
		eventZombiePicked := event.ZombiePicked{
			PickerGroupID: pickerGroupID,
			OutboxRecord:  record,
			Timestamp:     timestamp,
		}

		e(eventZombiePicked)
	}
}
