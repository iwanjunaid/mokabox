package event

import (
	"fmt"
	"time"
)

type RemoverStarted struct {
	PickerGroupID string
	Timestamp     time.Time
}

func (r RemoverStarted) String() string {
	return fmt.Sprintf("[%s:%s] Remover started", PREFIX, r.PickerGroupID)
}

func (r RemoverStarted) GetPickerGroupID() string {
	return r.PickerGroupID
}

func (r RemoverStarted) GetTimestamp() time.Time {
	return r.Timestamp
}
