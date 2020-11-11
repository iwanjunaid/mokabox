package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventRemoverStarted(e event.EventHandler, timestamp time.Time,
	pickerGroupID string) {
	if e != nil {
		eventRemoverStarted := event.RemoverStarted{
			PickerGroupID: pickerGroupID,
			Timestamp:     timestamp,
		}

		e(eventRemoverStarted)
	}
}
