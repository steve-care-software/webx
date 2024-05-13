package lists

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists/inserts"
)

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithInsert(insert inserts.Insert) Builder
	WithDelete(delete deletes.Delete) Builder
	Now() (List, error)
}

// List represents a list assignable
type List interface {
	IsInsert() bool
	Insert() inserts.Insert
	IsDelete() bool
	Delete() deletes.Delete
}
