package databases

import (
	"net/url"

	"github.com/steve-care-software/webx/domain/databases/blockchains"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/programs"
	"github.com/steve-care-software/webx/domain/databases/schemas"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewDatabaseBuilder creates a new database builder
func NewDatabaseBuilder() DatabaseBuilder {
	return createDatabaseBuilder()
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	publicKeyAdapter := keys.NewPublicKeyAdapter()
	return createContentBuilder(hashAdapter, publicKeyAdapter)
}

// NewMigrationBuilder creates a new migration builder
func NewMigrationBuilder() MigrationBuilder {
	hashAdapter := hash.NewAdapter()
	return createMigrationBuilder(hashAdapter)
}

// Builder represents a databases builder
type Builder interface {
	Create() Builder
	WithList(list []Database) Builder
	Now() (Databases, error)
}

// Databases represents databases
type Databases interface {
	List() []Database
}

// DatabaseBuilder represents a database builder
type DatabaseBuilder interface {
	Create() DatabaseBuilder
	WithContent(content Content) DatabaseBuilder
	WithConnections(connections []url.URL) DatabaseBuilder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Content() Content
	HasConnections() bool
	Connections() []url.URL
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithName(name string) ContentBuilder
	WithSchema(schema schemas.Schema) ContentBuilder
	WithReference(Reference blockchains.Blockchain) ContentBuilder
	WithMigration(migration Migration) ContentBuilder
	EncryptTo(encryptTo keys.PublicKey) ContentBuilder
	Now() (Content, error)
}

// Content represents the database content
type Content interface {
	Hash() hash.Hash
	Name() string
	Schema() schemas.Schema
	Reference() blockchains.Blockchain
	EncryptoTo() keys.PublicKey
	HasMigration() bool
	Migration() Migration
}

// MigrationBuilder represents a migration builder
type MigrationBuilder interface {
	Create() MigrationBuilder
	WithPrevious(previous Content) MigrationBuilder
	WithHeight(height uint) MigrationBuilder
	WithDescription(description string) MigrationBuilder
	WithProgram(program programs.Program) MigrationBuilder
	Now() (Migration, error)
}

// Migration represents a migration
type Migration interface {
	Hash() hash.Hash
	Previous() Content
	Height() uint
	Description() string
	HasProgram() bool
	Program() programs.Program
}
