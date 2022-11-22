package references

import (
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
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
)

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewKeysBuilder creates a new keys builder
func NewKeysBuilder() KeysBuilder {
	return createKeysBuilder()
}

// NewKeyBuilder creates a new key builder
func NewKeyBuilder() KeyBuilder {
	return createKeyBuilder()
}

// NewLinksBuilder creates a new links builder
func NewLinksBuilder() LinksBuilder {
	return createLinksBuilder()
}

// NewLinkBuilder creates a new link builder
func NewLinkBuilder() LinkBuilder {
	return createLinkBuilder()
}

// NewRelationsBuilder creates a new relations builder
func NewRelationsBuilder() RelationsBuilder {
	return createRelationsBuilder()
}

// NewRelationBuilder creates a new relation builder
func NewRelationBuilder() RelationBuilder {
	return createRelationBuilder()
}

// NewPointerAdapter creates a new pointer adapter
func NewPointerAdapter() PointerAdapter {
	builder := NewPointerBuilder()
	return createPointerAdapter(builder)
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Adapter represents a reference adapter
type Adapter interface {
	ToContent(ins Reference) ([]byte, error)
	ToReference(content []byte) (Reference, error)
}

// Factory represents a reference factory
type Factory interface {
	Create() (Reference, error)
}

// Builder represents a reference builder
type Builder interface {
	Create() Builder
	WithActive(active Keys) Builder
	WithPendings(pendings Keys) Builder
	WithDeleted(deleted Keys) Builder
	WithLinks(links Links) Builder
	WithRelations(relations Relations) Builder
	Now() (Reference, error)
}

// Reference represents the reference
type Reference interface {
	HasActive() bool
	Active() Keys
	HasPendings() bool
	Pendings() Keys
	HasDeleted() bool
	Deleted() Keys
	HasLinks() bool
	Links() Links
	HasRelations() bool
	Relations() Relations
}

// KeysAdapter represents the keys adapter
type KeysAdapter interface {
	ToContent(ins Keys) ([]byte, error)
	ToKeys(content []byte) (Keys, error)
}

// KeysBuilder represents a keys builder
type KeysBuilder interface {
	Create() KeysBuilder
	WithList(list []Key) KeysBuilder
	Now() (Keys, error)
}

// Keys represents keys
type Keys interface {
	List() []Key
	Fetch(hash hash.Hash) (Key, error)
}

// KeyAdapter represents the key adapter
type KeyAdapter interface {
	ToContent(ins Key) ([]byte, error)
	ToKey(content []byte) (Key, error)
}

// KeyBuilder represents a key builder
type KeyBuilder interface {
	Create() KeyBuilder
	WithHash(hash hash.Hash) KeyBuilder
	WithIndex(index uint) KeyBuilder
	WithKind(kind uint8) KeyBuilder
	WithContent(content Pointer) KeyBuilder
	WithTransaction(trx uint) KeyBuilder
	CreatedOn(createdOn time.Time) KeyBuilder
	IsEntity() KeyBuilder
	Now() (Key, error)
}

// Key represents a key
type Key interface {
	Hash() hash.Hash
	Index() uint
	Kind() uint8
	Content() Pointer
	IsEntity() bool
	CreatedOn() time.Time
	HasTransaction() bool
	Transaction() *uint
}

// LinksAdapter represents the links adapter
type LinksAdapter interface {
	ToContent(ins Links) ([]byte, error)
	ToLinks(content []byte) (Links, error)
}

// LinksBuilder represents a links builder
type LinksBuilder interface {
	Create() LinksBuilder
	WithList(list []Link) LinksBuilder
	Now() (Links, error)
}

// Links represents links
type Links interface {
	List() []Link
}

// LinkAdapter represents the link adapter
type LinkAdapter interface {
	ToContent(ins Link) ([]byte, error)
	ToLink(content []byte) (Link, error)
}

// LinkBuilder represents a link builder
type LinkBuilder interface {
	Create() LinkBuilder
	From(from uint) LinkBuilder
	To(to uint) LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	From() uint
	To() uint
}

// RelationsBuilder represents relations builder
type RelationsBuilder interface {
	Create() RelationsBuilder
	WithList(list []Relation) RelationsBuilder
	Now() (Relations, error)
}

// RelationsAdapter represents the relations adapter
type RelationsAdapter interface {
	ToContent(ins Relations) ([]byte, error)
	ToRelations(content []byte) (Relations, error)
}

// Relations represents relations
type Relations interface {
	List() []Relation
}

// RelationAdapter represents the relation adapter
type RelationAdapter interface {
	ToContent(ins Relation) ([]byte, error)
	ToRelation(content []byte) (Relation, error)
}

// RelationBuilder represents a relation builder
type RelationBuilder interface {
	Create() RelationBuilder
	From(from uint) RelationBuilder
	To(to []uint) RelationBuilder
	Now() (Relation, error)
}

// Relation represents a relation
type Relation interface {
	From() uint
	To() []uint
}

// PointerAdapter represents the pointer adapter
type PointerAdapter interface {
	ToContent(ins Pointer) ([]byte, error)
	ToPointer(content []byte) (Pointer, error)
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithLength(length uint) PointerBuilder
	From(from uint) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	From() uint
	Length() uint
}
