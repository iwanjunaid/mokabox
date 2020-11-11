package event

import (
	"fmt"
	"time"
)

type ZombiePickerStarted struct {
	PickerGroupID string
	Timestamp     time.Time
}

func (z ZombiePickerStarted) String() string {
	return fmt.Sprintf("[%s:%s] Zombie picker started", PREFIX, z.PickerGroupID)
}

func (z ZombiePickerStarted) GetPickerGroupID() string {
	return z.PickerGroupID
}

func (z ZombiePickerStarted) GetTimestamp() time.Time {
	return z.Timestamp
}
