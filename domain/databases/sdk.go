package databases

import (
	"net/url"

	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/databases/entities"
	"github.com/steve-care-software/webx/domain/databases/entries"
)

// Database represents a database
type Database interface {
	Content() Content
	HasConnections() bool
	Connections() []url.URL
}

// Content represents the database content
type Content interface {
	Name() string
	Blockchain() entities.Identifier
	Router() entities.Identifiers
	Sections() Sections
	HasEncryptTo() bool
	EncryptoTo() keys.PublicKey
	HasPrograms() bool
	Programs() Programs
	HasMigration() bool
	Migration() Migration
}

// Migration represents a migration
type Migration interface {
	Previous() Content
	Height() uint
	Description() string
}

// ProgramsBuilder represents the programs builder
type ProgramsBuilder interface {
	Create() ProgramsBuilder
	WithInit(init entities.Identifiers) ProgramsBuilder
	WithStop(stop entities.Identifiers) ProgramsBuilder
	WithStart(start entities.Identifiers) ProgramsBuilder
	WithDaemon(daemon entities.Identifiers) ProgramsBuilder
	Now() (Programs, error)
}

// Programs represents database programs
type Programs interface {
	HasInit() bool
	Init() entities.Identifiers
	HasStop() bool
	Stop() entities.Identifiers
	HasStart() bool
	Start() entities.Identifiers
	HasDaemon() bool
	Daemon() entities.Identifiers
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
