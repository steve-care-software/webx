package databases

import (
	"net/url"
	"time"

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

	// KindApplication represents an application kind
	KindApplication

	// KindEntry represents an entry kind
	KindEntry
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

// NewSectionsBuilder creates a new sections builder
func NewSectionsBuilder() SectionsBuilder {
	return createSectionsBuilder()
}

// NewSectionBuilder creates a new section builder
func NewSectionBuilder() SectionBuilder {
	return createSectionBuilder()
}

// NewPointersBuilder creates a new pointers builder
func NewPointersBuilder() PointersBuilder {
	return createPointersBuilder()
}

// NewPointerBulder creates a new pointer builder
func NewPointerBulder() PointerBuilder {
	return createPointerBuilder()
}

// NewSizeInBytesBuilder creates a new size in bytes builder
func NewSizeInBytesBuilder() SizeInBytesBuilder {
	return createSizeInBytesBuilder()
}

// NewEntriesBuilder creates a new entries builder
func NewEntriesBuilder() EntriesBuilder {
	return createEntriesBuilder()
}

// NewEntryBuilder creates a new entry builder
func NewEntryBuilder() EntryBuilder {
	return createEntryBuilder()
}

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithHead(head Head) Builder
	WithPendings(pendings Entries) Builder
	WithConnections(connections []url.URL) Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Head() Head
	HasPendings() bool
	Pendings() Entries
	HasConnections() bool
	Connections() []url.URL
}

// HeadBuilder represents a head builder
type HeadBuilder interface {
	Create() HeadBuilder
	WithName(name string) HeadBuilder
	WithSections(sections Sections) HeadBuilder
	WithBlockInterval(blockInterval time.Duration) HeadBuilder
	WithSyncInterval(syncInterval time.Duration) HeadBuilder
	WithMigration(migration Migration) HeadBuilder
	Now() (Head, error)
}

// Head represents the database head
type Head interface {
	Name() string
	Sections() Sections
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

// SectionsBuilder represents a sections builder
type SectionsBuilder interface {
	Create() SectionsBuilder
	WithList(list []Section) SectionsBuilder
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
	WithKind(kind uint8) SectionBuilder
	WithPointers(pointers Pointers) SectionBuilder
	Now() (Section, error)
}

// Section represents a section
type Section interface {
	Index() uint
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
	CreatedOn(createdOn time.Time) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	BeginsOn() SizeInBytes
	Length() SizeInBytes
	CreatedOn() time.Time
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
	IsZero() bool
}

// EntriesBuilder represents an entries builder
type EntriesBuilder interface {
	Create() EntriesBuilder
	WithList(list []Entry) EntriesBuilder
	Now() (Entries, error)
}

// Entries represents entries
type Entries interface {
	List() []Entry
}

// EntryBuilder represents an entry builder
type EntryBuilder interface {
	Create() EntryBuilder
	WithEntity(entity entities.Entity) EntryBuilder
	WithTransaction(transaction transactions.Transaction) EntryBuilder
	WithPointer(pointer Pointer) EntryBuilder
	WithContent(content []byte) EntryBuilder
	WithKind(kind uint8) EntryBuilder
	Now() (Entry, error)
}

// Entry represents a database entry
type Entry interface {
	Entity() entities.Entity
	Pointer() Pointer
	Content() []byte
	Kind() uint8
	HasTransaction() bool
	Transaction() transactions.Transaction
}
