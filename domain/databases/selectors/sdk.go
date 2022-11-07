package selectors

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithToken(token entities.Identifier) Builder
	WithInside(inside entities.Identifier) Builder
	WithFunc(fn entities.Identifier) Builder
	Now() (Selector, error)
}

// Selector represents a selector
type Selector interface {
	Entity() entities.Entity
	Token() entities.Identifier
	Inside() entities.Identifier
	Func() entities.Identifier
}
