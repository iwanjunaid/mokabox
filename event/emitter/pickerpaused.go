package emitter

import (
	"time"

	"github.com/google/uuid"
	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventPickerPaused(e event.EventHandler, timestamp time.Time,
	pickerGroupID uuid.UUID) {
	if e != nil {
		eventPickerPaused := event.PickerPaused{
			PickerGroupID: pickerGroupID,
			Timestamp:     timestamp,
		}

		e(eventPickerPaused)
	}
}
