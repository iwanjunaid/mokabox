package event

import (
	"github.com/iwanjunaid/mokabox/internal/interfaces/event"
)

const PREFIX = "mokabox"

type EventHandler func(event.Event)
