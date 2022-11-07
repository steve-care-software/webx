package databases

import (
	"net/url"
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/databases/blockchains/transactions"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

const (
	// KindBlockchain represents a blockchain kind
	KindBlockchain uint8 = iota

	// KindBlockchainBlock represents the blockchain's block kind
	KindBlockchainBlock

	// KindBlockchainTransaction represents the blockchain's transaction kind
	KindBlockchainTransaction

	// KindIdentity represents an identity kind
	KindIdentity

	// KindIdentityModification represents the identity's modification kind
	KindIdentityModification

	// KindGrammar represents a grammar kind
	KindGrammar

	// KindGrammarCardinality represents the grammar's cardinality kind
	KindGrammarCardinality

	// KindGrammarChannel represents the grammar's channel kind
	KindGrammarChannel

	// KindGrammarElement represents the grammar's element kind
	KindGrammarElement

	// KindGrammarEverything represents the grammar's everything kind
	KindGrammarEverything

	// KindGrammarLine represents the grammar's line kind
	KindGrammarLine

	// KindGrammarSuite represents the grammar's suite kind
	KindGrammarSuite

	// KindGrammarToken represents the grammar's token kind
	KindGrammarToken

	// KindSelector represents a selector kind
	KindSelector

	// KindSelectorFetcher represents the selector's fetcher kind
	KindSelectorFetcher

	// KindSelectorFunc represents the selector's func kind
	KindSelectorFunc

	// KindSelectorInside represents the selector's inside kind
	KindSelectorInside

	// KindSelectorToken represents the selector's token kind
	KindSelectorToken

	// KindTree represents a tree kind
	KindTree

	// KindTreeContent represents the tree's content kind
	KindTreeContent

	// KindTreeElement represents the tree's element kind
	KindTreeElement

	// KindTreeLine represents the tree's line kind
	KindTreeLine

	// KindProgram represents a program kind
	KindProgram

	// KindProgramApplication represents the program's application kind
	KindProgramApplication

	// KindProgramInstruction represents the program's instruction kind
	KindProgramInstruction

	// KindProgramValue represents the program's value kind
	KindProgramValue

	// KindRoute represents a route kind
	KindRoute

	// KindEntry represents an entry kind
	KindEntry
)

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithConnections(connections []url.URL) Builder
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
	WithSections(sections Section) ContentBuilder
	WithEncrypTo(encryptTo keys.PublicKey) ContentBuilder
	WithBlockchain(blockchain entities.Identifier) ContentBuilder
	WithRoutes(routes entities.Identifiers) ContentBuilder
	WithPrograms(programs Programs) ContentBuilder
	WithMigration(migration Migration) ContentBuilder
	Now() (Content, error)
}

// Content represents the database content
type Content interface {
	Name() string
	Sections() Sections
	EncryptoTo() keys.PublicKey
	HasBlockchain() bool
	Blockchain() entities.Identifier
	HasRoutes() bool
	Routes() entities.Identifiers
	HasPrograms() bool
	Programs() Programs
	HasMigration() bool
	Migration() Migration
}

// MigrationBuilder represents a migration builder
type MigrationBuilder interface {
	Create() MigrationBuilder
	WithPrevious(previous Content) MigrationBuilder
	WithHeight(height uint) MigrationBuilder
	WithDescrition(description string) MigrationBuilder
	Now() (Migration, error)
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

// SectionsBuilder represents a sections builder
type SectionsBuilder interface {
	Create() SectionsBuilder
	WithList(list []Pointer) SectionsBuilder
	Now() (Sections, error)
}

// Sections represents sections
type Sections interface {
	List() []Section
}

// SectionBuilder represents a section builder
type SectionBuilder interface {
	Create() SectionBuilder
	WithIndex(index uint) SectionBuilder
	WithLength(length uint) SectionBuilder
	WithKind(kind uint8) SectionBuilder
	WithPointers(pointers Pointers) SectionBuilder
	Now() (Section, error)
}

// Section represents a section
type Section interface {
	Index() uint
	Length() uint
	Kind() uint8
	HasPointers() bool
	Pointers() Pointers
}

// PointersBuilder represents a pointers builder
type PointersBuilder interface {
	Create() PointersBuilder
	WithList(list []Pointer) PointersBuilder
	Now() (Pointers, error)
}

// Pointers represents pointers
type Pointers interface {
	List() []Pointer
}

// PointerBuilder represents the pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithLength(length SizeInBytes) PointerBuilder
	WithReferences(references Pointers) PointerBuilder
	BeginsOn(beginsOn SizeInBytes) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	BeginsOn() SizeInBytes
	Length() SizeInBytes
	HasReferences() bool
	References() Pointers
}

// SizeInBytesBuilder represents a size in bytes builder
type SizeInBytesBuilder interface {
	Create() SizeInBytesBuilder
	WithMaxAmount(maxAmount uint) SizeInBytesBuilder
	WithAmount(amount uint) SizeInBytesBuilder
	Now() (SizeInBytes, error)
}

// SizeInBytes represents a sizeInBytes
type SizeInBytes interface {
	MaxAmount() uint
	Amount() uint
}

// EntryBuilder represents an entry builder
type EntryBuilder interface {
	Create() EntryBuilder
	WithEntity(entity entities.Entity) EntryBuilder
	WithTransaction(transaction transactions.Transaction) EntryBuilder
	WithPointer(pointer Pointer) EntryBuilder
	WithContent(content []byte) EntryBuilder
	WithKind(kind uint8) EntryBuilder
	CreatedOn(createdOn time.Time) EntryBuilder
	Now() (Entry, error)
}

// Entry represents a database entry
type Entry interface {
	Entity() entities.Entity
	Transaction() transactions.Transaction
	Pointer() Pointer
	Content() []byte
	Kind() uint8
	CreatedOn() time.Time
}
