package event

import (
	"fmt"
	"time"
)

type PickerStarted struct {
	PickerGroupID string
	Timestamp     time.Time
}

func (p PickerStarted) String() string {
	return fmt.Sprintf("[%s:%s] Picker started", PREFIX, p.PickerGroupID)
}

func (p PickerStarted) GetPickerGroupID() string {
	return p.PickerGroupID
}

func (p PickerStarted) GetTimestamp() time.Time {
	return p.Timestamp
}
