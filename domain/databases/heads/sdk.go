package heads

import (
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

// Builder represents a head builder
type Builder interface {
	Create() Builder
	WithActive(active Keys) Builder
	WithDeleted(deleted Keys) Builder
	WithLinks(links Links) Builder
	WithRelations(relations Relations) Builder
	WithWeightedRelations(weightedRelations WeightedRelations) Builder
	WithBodyIndex(bodyIndex uint) Builder
	Now() (Head, error)
}

// Head represents the head
type Head interface {
	Active() Keys
	Deleted() Keys
	Links() Links
	Relations() Relations
	WeightedRelations() WeightedRelations
	BodyIndex() uint
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
	Deleted() []Key
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
	Transaction() uint
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
	Now() (WeightedElement, error)
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
