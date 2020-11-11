package event

import (
	"fmt"
	"time"

	"github.com/iwanjunaid/mokabox/model"
)

type ZombieAcquired struct {
	PickerGroupID string
	OriginGroupID string
	OutboxRecord  *model.OutboxRecord
	Timestamp     time.Time
}

func (z ZombieAcquired) String() string {
	id := z.OutboxRecord.ID
	originGroupID := z.OriginGroupID
	groupID := z.OutboxRecord.GroupID

	return fmt.Sprintf("[%s:%s] Zombie message with ID %s from Origin Group ID %s is successfully acquired",
		PREFIX, groupID, id.Hex(), originGroupID)
}

func (z ZombieAcquired) GetPickerGroupID() string {
	return z.PickerGroupID
}

func (z ZombieAcquired) GetOriginGroupID() string {
	return z.OriginGroupID
}

func (z ZombieAcquired) GetRecord() *model.OutboxRecord {
	return z.OutboxRecord
}

func (z ZombieAcquired) GetTimestamp() time.Time {
	return z.Timestamp
}
