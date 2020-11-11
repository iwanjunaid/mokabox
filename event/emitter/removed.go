package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
	"github.com/iwanjunaid/mokabox/model"
)

func EmitEventRemoved(e event.EventHandler, timestamp time.Time,
	pickerGroupID string, record *model.OutboxRecord) {
	if e != nil {
		eventRemoved := event.Removed{
			PickerGroupID: pickerGroupID,
			OutboxRecord:  record,
			Timestamp:     timestamp,
		}

		e(eventRemoved)
	}
}
