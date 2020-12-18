package config

type OutboxConfig interface {
	GetGroupID() string
	GetDatabaseName() string
	GetOutboxCollectionName() string
	GetPickerPollInterval() int
	GetPickerMessageLimitPerPoll() int
	GetZombieInterval() int
	GetZombiePickerPollInterval() int
	GetZombiePickerMessageLimitPerPoll() int
	GetRemoverPollInterval() int
	GetRemoverMessageLimitPerPoll() int
}
