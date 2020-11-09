package event

import (
	"github.com/iwanjunaid/mokabox/internal/interfaces/event"
)

const PREFIX = "pokabox"

type EventHandler func(event.Event)
