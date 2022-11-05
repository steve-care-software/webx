package databases

import (
	"net/url"

	"github.com/steve-care-software/webx/domain/blockchains"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/databases/entities"
	"github.com/steve-care-software/webx/domain/databases/entries"
)

// Database represents a database
type Database interface {
	Name() string
	Content() Content
	HasConnections() bool
	Connections() []url.URL
}

// Content represents the database content
type Content interface {
	Name() []byte
	Sections() Sections
	Reference() blockchains.Blockchain
	HasEncryptTo() bool
	EncryptoTo() keys.PublicKey
	HasMigration() bool
	Migration() Migration
}

// Migration represents a migration
type Migration interface {
	Previous() Content
	Height() uint
	Description() string
	HasProgram() bool
	Program() entities.Identifier
}

// Sections represents sections
type Sections interface {
	List() []Section
}

// Section represents a section
type Section interface {
	Index() uint
	Length() uint
	Kind() entries.Kind
	Pointers() Pointers
}

// Pointers represents pointers
type Pointers interface {
	List() []Pointer
}

// Pointer represents a pointer
type Pointer interface {
	Index() uint
	BeginsOn() []uint
	Length() uint
}
