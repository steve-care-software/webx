package selectors

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithGrammar(grammar entities.Identifier) Builder
	WithToken(token entities.Identifier) Builder
	WithInside(inside entities.Identifier) Builder
	WithFunc(fn entities.Identifier) Builder
	Now() (Selector, error)
}

// Selector represents a selector
type Selector interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Token() entities.Identifier
	Inside() entities.Identifier
	Func() entities.Identifier
}
