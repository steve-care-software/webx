package databases

import (
	"time"

	"github.com/steve-care-software/webx/domain/databases/references"
)

type head struct {
	name          string
	reference     references.Reference
	blockInterval time.Duration
	syncInterval  time.Duration
	migration     Migration
}

func createHead(
	name string,
	reference references.Reference,
	blockInterval time.Duration,
	syncInterval time.Duration,
) Head {
	return createHeadInternally(name, reference, blockInterval, syncInterval, nil)
}

func createHeadWithMigration(
	name string,
	reference references.Reference,
	blockInterval time.Duration,
	syncInterval time.Duration,
	migration Migration,
) Head {
	return createHeadInternally(name, reference, blockInterval, syncInterval, migration)
}

func createHeadInternally(
	name string,
	reference references.Reference,
	blockInterval time.Duration,
	syncInterval time.Duration,
	migration Migration,
) Head {
	out := head{
		name:          name,
		reference:     reference,
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

// Reference returns the reference
func (obj *head) Reference() references.Reference {
	return obj.reference
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
