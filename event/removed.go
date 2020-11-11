package event

import (
	"fmt"
	"time"

	"github.com/iwanjunaid/mokabox/model"
)

type Removed struct {
	PickerGroupID string
	OutboxRecord  *model.OutboxRecord
	Timestamp     time.Time
}

func (r Removed) String() string {
	return fmt.Sprintf("[%s:%s] Message with ID %s successfully removed",
		PREFIX, r.PickerGroupID, r.OutboxRecord.ID.Hex())
}

func (r Removed) GetPickerGroupID() string {
	return r.PickerGroupID
}

func (r Removed) GetRecord() *model.OutboxRecord {
	return r.OutboxRecord
}

func (r Removed) GetTimestamp() time.Time {
	return r.Timestamp
}
