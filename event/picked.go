package event

import (
	"fmt"
	"time"

	"github.com/iwanjunaid/mokabox/model"
)

type Picked struct {
	OutboxRecord *model.OutboxRecord
	Timestamp    time.Time
}

func (f Picked) String() string {
	id := f.OutboxRecord.ID.Hex()
	groupID := f.OutboxRecord.GroupID

	return fmt.Sprintf("[%s:%s] Message with ID %s picked", PREFIX, groupID, id)
}

func (p Picked) GetPickerGroupID() string {
	return p.OutboxRecord.GroupID
}

func (p Picked) GetRecord() *model.OutboxRecord {
	return p.OutboxRecord
}

func (p Picked) GetTimestamp() time.Time {
	return p.Timestamp
}
