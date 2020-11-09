package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
	"github.com/iwanjunaid/mokabox/model"
)

func EmitEventPicked(e event.EventHandler, timestamp time.Time, record *model.OutboxRecord) {
	if e != nil {
		eventPicked := event.Picked{
			OutboxRecord: record,
			Timestamp:    timestamp,
		}

		e(eventPicked)
	}
}
