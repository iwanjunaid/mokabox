package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventRemoverPaused(e event.EventHandler, timestamp time.Time,
	pickerGroupID string) {
	if e != nil {
		eventRemoverPaused := event.RemoverPaused{
			PickerGroupID: pickerGroupID,
			Timestamp:     timestamp,
		}

		e(eventRemoverPaused)
	}
}
