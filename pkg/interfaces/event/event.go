package event

import (
	"fmt"
	"time"
)

type Event interface {
	fmt.Stringer
	GetPickerGroupID() string
	GetTimestamp() time.Time
}
