package databases

import "time"

type head struct {
	name          string
	sections      Sections
	blockInterval time.Duration
	syncInterval  time.Duration
	migration     Migration
}

func createHead(
	name string,
	sections Sections,
	blockInterval time.Duration,
	syncInterval time.Duration,
) Head {
	return createHeadInternally(name, sections, blockInterval, syncInterval, nil)
}

func createHeadWithMigration(
	name string,
	sections Sections,
	blockInterval time.Duration,
	syncInterval time.Duration,
	migration Migration,
) Head {
	return createHeadInternally(name, sections, blockInterval, syncInterval, migration)
}

func createHeadInternally(
	name string,
	sections Sections,
	blockInterval time.Duration,
	syncInterval time.Duration,
	migration Migration,
) Head {
	out := head{
		name:          name,
		sections:      sections,
		blockInterval: blockInterval,
		syncInterval:  syncInterval,
		migration:     migration,
	}

	return &out
}

// Name returns the name
func (obj *head) Name() string {
	return obj.name
}

// Sections returns the sections
func (obj *head) Sections() Sections {
	return obj.sections
}

// BlockInterval returns the block interval
func (obj *head) BlockInterval() time.Duration {
	return obj.blockInterval
}

// SyncInterval returns the sync interval
func (obj *head) SyncInterval() time.Duration {
	return obj.syncInterval
}

// HasMigration returns true if there is migration, false otherwise
func (obj *head) HasMigration() bool {
	return obj.migration != nil
}

// Migration returns the migration, if any
func (obj *head) Migration() Migration {
	return obj.migration
}
