package operators

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/integers"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/relationals"
)

// NewBuilder creates a new operator builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an operator builder
type Builder interface {
	Create() Builder
	WithRelational(relational relationals.Relational) Builder
	WithInteger(integer integers.Integer) Builder
	IsEqual() Builder
	Now() (Operator, error)
}

// Operator represents an operator
type Operator interface {
	Hash() hash.Hash
	IsEqual() bool
	IsRelational() bool
	Relational() relationals.Relational
	IsInteger() bool
	Integer() integers.Integer
}
