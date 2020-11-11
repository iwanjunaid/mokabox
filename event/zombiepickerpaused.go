package event

import (
	"fmt"
	"time"
)

type ZombiePickerPaused struct {
	PickerGroupID string
	Timestamp     time.Time
}

func (z ZombiePickerPaused) String() string {
	return fmt.Sprintf("[%s:%s] Zombie picker paused", PREFIX, z.PickerGroupID)
}

func (z ZombiePickerPaused) GetPickerGroupID() string {
	return z.PickerGroupID
}

func (z ZombiePickerPaused) GetTimestamp() time.Time {
	return z.Timestamp
}
