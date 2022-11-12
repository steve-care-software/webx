package heads

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

	// KindIdentityName represents an identity name kind
	KindIdentityName

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

// NewWeightedRelationsBuilder creates a new weightedRelations builder
func NewWeightedRelationsBuilder() WeightedRelationsBuilder {
	return createWeightedRelationsBuilder()
}

// NewWeightedRelationBuilder creates a new weightedRelation builder
func NewWeightedRelationBuilder() WeightedRelationBuilder {
	return createWeightedRelationBuilder()
}

// NewWeightedElementsBuiler creates a new weighted elements builder
func NewWeightedElementsBuiler() WeightedElementsBuilder {
	return createWeightedElementsBuilder()
}

// NewWeightedElementBuiler creates a new weighted element builder
func NewWeightedElementBuiler() WeightedElementBuilder {
	return createWeightedElementBuilder()
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Builder represents a head builder
type Builder interface {
	Create() Builder
	WithActive(active Keys) Builder
	WithDeleted(deleted Keys) Builder
	WithLinks(links Links) Builder
	WithRelations(relations Relations) Builder
	WithWeightedRelations(weightedRelations WeightedRelations) Builder
	Now() (Head, error)
}

// Head represents the head
type Head interface {
	HasActive() bool
	Active() Keys
	HasDeleted() bool
	Deleted() Keys
	HasLinks() bool
	Links() Links
	HasRelations() bool
	Relations() Relations
	HasWeightedRelations() bool
	WeightedRelations() WeightedRelations
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

// Relations represents relations
type Relations interface {
	List() []Relation
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

// WeightedRelationsBuilder represents weighted relations builder
type WeightedRelationsBuilder interface {
	Create() WeightedRelationsBuilder
	WithList(list []WeightedRelation) WeightedRelationsBuilder
	Now() (WeightedRelations, error)
}

// WeightedRelations represents weighted relations
type WeightedRelations interface {
	List() []WeightedRelation
}

// WeightedRelationBuilder represents a weighted relation builder
type WeightedRelationBuilder interface {
	Create() WeightedRelationBuilder
	From(from uint) WeightedRelationBuilder
	To(to WeightedElements) WeightedRelationBuilder
	Now() (WeightedRelation, error)
}

// WeightedRelation represents a weighted relation
type WeightedRelation interface {
	From() uint
	To() WeightedElements
}

// WeightedElementsBuilder represents weighted elements builder
type WeightedElementsBuilder interface {
	Create() WeightedElementsBuilder
	WithList(list []WeightedElement) WeightedElementsBuilder
	Now() (WeightedElements, error)
}

// WeightedElements represents weighted elements
type WeightedElements interface {
	List() []WeightedElement
}

// WeightedElementBuilder represents a weighted element builder
type WeightedElementBuilder interface {
	Create() WeightedElementBuilder
	WithIndex(index uint) WeightedElementBuilder
	WithWeight(weight uint) WeightedElementBuilder
	Now() (WeightedElement, error)
}

// WeightedElement represents a weighted element
type WeightedElement interface {
	Index() uint
	Weight() uint
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
