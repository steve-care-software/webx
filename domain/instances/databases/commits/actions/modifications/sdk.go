package modifications

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/updates"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewModificationBuilder creates a new modification builder
func NewModificationBuilder() ModificationBuilder {
	hashAdapter := hash.NewAdapter()
	return createModificationBuilder(
		hashAdapter,
	)
}

// Builder represents a modifications builder
type Builder interface {
	Create() Builder
	WithList(list []Modification) Builder
	Now() (Modifications, error)
}

// Modifications represents modifications
type Modifications interface {
	Hash() hash.Hash
	List() []Modification
}

// ModificationBuilder represents a modification builder
type ModificationBuilder interface {
	Create() ModificationBuilder
	WithInsert(insert []byte) ModificationBuilder
	WithUpdate(update updates.Update) ModificationBuilder
	WithDelete(del deletes.Delete) ModificationBuilder
	Now() (Modification, error)
}

// Modification represents a modification
type Modification interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() []byte
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() deletes.Delete
}
