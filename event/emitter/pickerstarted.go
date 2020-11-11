package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventPickerStarted(e event.EventHandler, timestamp time.Time,
	pickerGroupID string) {
	if e != nil {
		pickerStarted := event.PickerStarted{
			PickerGroupID: pickerGroupID,
			Timestamp:     timestamp,
		}

		e(pickerStarted)
	}
}
