package config

import (
	"strings"

	"github.com/iwanjunaid/mokabox/pkg/interfaces/config"
)

type CommonOutboxConfig struct {
	GroupID                         string
	DatabaseName                    string
	OutboxCollectionName            string
	PickerPollInterval              int
	PickerMessageLimitPerPoll       int
	ZombieInterval                  int
	ZombiePickerPollInterval        int
	ZombiePickerMessageLimitPerPoll int
	RemoverPollInterval             int
	RemoverMessageLimitPerPoll      int
}

func NewDefaultCommonOutboxConfig(groupID string, databaseName string) config.OutboxConfig {
	return NewCommonOutboxConfig(groupID, databaseName, "", 0, 0, 0, 0, 0, 0, 0)
}

func NewCommonOutboxConfig(groupID string, databaseName string,
	outboxCollectionName string,
	pickerPollInterval int, pickerMessageLimitPerPoll int,
	zombieInterval int, zombiePickerPollInterval int, zombiePickerMessageLimitPerPoll int,
	removerPollInterval int, removerMessageLimitPerPoll int) config.OutboxConfig {

	cOutboxCollectionName := "outbox"

	if trimmed := strings.TrimSpace(outboxCollectionName); len(trimmed) > 0 {
		cOutboxCollectionName = trimmed
	}

	cPickerPollInterval := 3

	if pickerPollInterval >= 1 {
		cPickerPollInterval = pickerPollInterval
	}

	cPickerMessageLimitPerPoll := 100

	if pickerMessageLimitPerPoll >= 1 {
		cPickerMessageLimitPerPoll = pickerMessageLimitPerPoll
	}

	cZombieInterval := 60 * 3

	if zombieInterval >= 60*3 {
		cZombieInterval = zombieInterval
	}

	cZombiePickerPollInterval := 3

	if zombiePickerPollInterval >= 1 {
		cZombiePickerPollInterval = zombiePickerPollInterval
	}

	cZombiePickerMessageLimitPerPoll := 100

	if zombiePickerMessageLimitPerPoll >= 1 {
		cZombiePickerMessageLimitPerPoll = zombiePickerMessageLimitPerPoll
	}

	cRemoverPollInterval := 3

	if removerPollInterval >= 1 {
		cRemoverPollInterval = removerPollInterval
	}

	cRemoverMessageLimitPerPoll := 100

	if removerMessageLimitPerPoll >= 1 {
		cRemoverMessageLimitPerPoll = removerMessageLimitPerPoll
	}

	config := &CommonOutboxConfig{
		groupID,
		databaseName,
		cOutboxCollectionName,
		cPickerPollInterval,
		cPickerMessageLimitPerPoll,
		cZombieInterval,
		cZombiePickerPollInterval,
		cZombiePickerMessageLimitPerPoll,
		cRemoverPollInterval,
		cRemoverMessageLimitPerPoll,
	}

	return config
}

func (c *CommonOutboxConfig) GetGroupID() string {
	return c.GroupID
}

func (c *CommonOutboxConfig) GetDatabaseName() string {
	return c.DatabaseName
}

func (c *CommonOutboxConfig) GetOutboxCollectionName() string {
	return c.OutboxCollectionName
}

func (c *CommonOutboxConfig) GetPickerPollInterval() int {
	return c.PickerPollInterval
}

func (c *CommonOutboxConfig) GetPickerMessageLimitPerPoll() int {
	return c.PickerMessageLimitPerPoll
}

func (c *CommonOutboxConfig) GetZombieInterval() int {
	return c.ZombieInterval
}

func (c *CommonOutboxConfig) GetZombiePickerPollInterval() int {
	return c.ZombiePickerPollInterval
}

func (c *CommonOutboxConfig) GetZombiePickerMessageLimitPerPoll() int {
	return c.ZombiePickerMessageLimitPerPoll
}

func (c *CommonOutboxConfig) GetRemoverPollInterval() int {
	return c.RemoverPollInterval
}

func (c *CommonOutboxConfig) GetRemoverMessageLimitPerPoll() int {
	return c.RemoverMessageLimitPerPoll
}
