package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventZombiePickerPaused(e event.EventHandler, timestamp time.Time,
	pickerGroupID string) {
	if e != nil {
		eventZombiePickerPaused := event.ZombiePickerPaused{
			PickerGroupID: pickerGroupID,
			Timestamp:     timestamp,
		}

		e(eventZombiePickerPaused)
	}
}
