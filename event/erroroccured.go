package event

import (
	"fmt"
	"time"
)

type ErrorOccured struct {
	PickerGroupID string
	Error         error
	Timestamp     time.Time
}

func (e ErrorOccured) String() string {
	return fmt.Sprintf("[%s:%s] Error occured: %s", PREFIX, e.PickerGroupID, e.Error.Error())
}

func (e ErrorOccured) GetPickerGroupID() string {
	return e.PickerGroupID
}

func (e ErrorOccured) GetError() error {
	return e.Error
}

func (e ErrorOccured) GetTimestamp() time.Time {
	return e.Timestamp
}
