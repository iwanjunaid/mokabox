package emitter

import (
	"time"

	"github.com/iwanjunaid/mokabox/event"
	"github.com/iwanjunaid/mokabox/model"
)

func EmitEventStatusChanged(e event.EventHandler, timestamp time.Time,
	fromStatus string, toStatus string, record *model.OutboxRecord) {
	if e != nil {
		eventStatusChanged := event.StatusChanged{
			From:         fromStatus,
			To:           toStatus,
			OutboxRecord: record,
			Timestamp:    timestamp,
		}

		e(eventStatusChanged)
	}
}
