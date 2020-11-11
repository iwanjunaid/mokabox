package event

import (
	"fmt"
	"time"

	"github.com/iwanjunaid/mokabox/model"
)

type ZombiePicked struct {
	PickerGroupID string
	OutboxRecord  *model.OutboxRecord
	Timestamp     time.Time
}

func (z ZombiePicked) String() string {
	pickerGroupID := z.PickerGroupID
	id := z.OutboxRecord.ID
	originGroupID := z.OutboxRecord.GroupID

	return fmt.Sprintf("[%s:%s] Zombie message with ID %s from Origin Group ID %s is picked",
		PREFIX, pickerGroupID, id.Hex(), originGroupID)
}

func (z ZombiePicked) GetPickerGroupID() string {
	return z.PickerGroupID
}

func (z ZombiePicked) GetRecord() *model.OutboxRecord {
	return z.OutboxRecord
}

func (z ZombiePicked) GetTimestamp() time.Time {
	return z.Timestamp
}
