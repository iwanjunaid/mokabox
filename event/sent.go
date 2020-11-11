package event

import (
	"fmt"
	"time"

	"github.com/iwanjunaid/mokabox/model"
)

type Sent struct {
	PickerGroupID  string
	KafkaTopic     string
	KafkaPartition int32
	OutboxRecord   *model.OutboxRecord
	Timestamp      time.Time
}

func (s Sent) String() string {
	return fmt.Sprintf("[%s:%s] Message with ID %s sent to kafka in topic %s and partition %d",
		PREFIX, s.PickerGroupID, s.OutboxRecord.ID.Hex(),
		s.KafkaTopic, s.KafkaPartition)
}

func (s Sent) GetPickerGroupID() string {
	return s.PickerGroupID
}

func (s Sent) GetRecord() *model.OutboxRecord {
	return s.OutboxRecord
}

func (s Sent) GetTimestamp() time.Time {
	return s.Timestamp
}
