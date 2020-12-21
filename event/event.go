package event

import (
	"github.com/iwanjunaid/mokabox/pkg/interfaces/event"
)

const PREFIX = "mokabox"

type EventHandler func(event.Event)
