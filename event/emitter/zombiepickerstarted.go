package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
)

func EmitEventZombiePickerStarted(e event.EventHandler, timestamp time.Time,
	pickerGroupID string) {
	if e != nil {
		eventZombiePickerStarted := event.ZombiePickerStarted{
			PickerGroupID: pickerGroupID,
			Timestamp:     timestamp,
		}

		e(eventZombiePickerStarted)
	}
}
