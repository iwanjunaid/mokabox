package event

import (
	"fmt"
	"time"
)

type PickerPaused struct {
	PickerGroupID string
	Timestamp     time.Time
}

func (p PickerPaused) String() string {
	return fmt.Sprintf("[%s:%s] Picker paused", PREFIX, p.PickerGroupID)
}

func (p PickerPaused) GetPickerGroupID() string {
	return p.PickerGroupID
}

func (p PickerPaused) GetTimestamp() time.Time {
	return p.Timestamp
}
