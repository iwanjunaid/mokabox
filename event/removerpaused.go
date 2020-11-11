package event

import (
	"fmt"
	"time"
)

type RemoverPaused struct {
	PickerGroupID string
	Timestamp     time.Time
}

func (r RemoverPaused) String() string {
	return fmt.Sprintf("[%s:%s] Remover paused", PREFIX, r.PickerGroupID)
}

func (r RemoverPaused) GetPickerGroupID() string {
	return r.PickerGroupID
}

func (r RemoverPaused) GetTimestamp() time.Time {
	return r.Timestamp
}
