package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventErrorOccured(e event.EventHandler, timestamp time.Time,
	pickerGroupID string, err error) {
	if e != nil {
		eventErrorOccured := event.ErrorOccured{
			PickerGroupID: pickerGroupID,
			Error:         err,
			Timestamp:     timestamp,
		}

		e(eventErrorOccured)
	}
}
