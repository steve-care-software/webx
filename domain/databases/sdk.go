package databases

import (
	"net/url"
	"time"

	"github.com/steve-care-software/webx/domain/databases/references"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewHeadBuilder creates a new head builder
func NewHeadBuilder() HeadBuilder {
	return createHeadBuilder()
}

// NewMigrationBuilder creates a new migration builder
func NewMigrationBuilder() MigrationBuilder {
	return createMigrationBuilder()
}

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithHead(head Head) Builder
	WithConnections(connections []url.URL) Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Head() Head
	HasConnections() bool
	Connections() []url.URL
}

// HeadBuilder represents a head builder
type HeadBuilder interface {
	Create() HeadBuilder
	WithName(name string) HeadBuilder
	WithReference(reference references.Reference) HeadBuilder
	WithBlockInterval(blockInterval time.Duration) HeadBuilder
	WithSyncInterval(syncInterval time.Duration) HeadBuilder
	WithMigration(migration Migration) HeadBuilder
	Now() (Head, error)
}

// Head represents the database head
type Head interface {
	Name() string
	Reference() references.Reference
	BlockInterval() time.Duration
	SyncInterval() time.Duration
	HasMigration() bool
	Migration() Migration
}

// MigrationBuilder represents a migration builder
type MigrationBuilder interface {
	Create() MigrationBuilder
	WithPrevious(previous Head) MigrationBuilder
	WithHeight(height uint) MigrationBuilder
	WithDescrition(description string) MigrationBuilder
	Now() (Migration, error)
}

// Migration represents a migration
type Migration interface {
	Previous() Head
	Height() uint
	Description() string
}
