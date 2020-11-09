package emitter

import (
	"time"

	"github.com/google/uuid"
	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventPickerStarted(e event.EventHandler, timestamp time.Time, pickerGroupID uuid.UUID) {
	if e != nil {
		pickerStarted := event.PickerStarted{
			PickerGroupID: pickerGroupID,
			Timestamp:     timestamp,
		}

		e(pickerStarted)
	}
}
