package routes

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a route builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithGrammar(grammar entities.Identifier) Builder
	WithSelectors(selectors entities.Identifiers) Builder
	WithProgram(program entities.Identifier) Builder
	Now() (Route, error)
}

// Route represents a route
type Route interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Selectors() entities.Identifiers
	Program() entities.Identifier
}
