package tokens

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a token builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithLines(lines entities.Identifiers) Builder
	WithSuites(suites entities.Identifiers) Builder
	Now() (Token, error)
}

// Token represents token metadata
type Token interface {
	Entity() entities.Entity
	Lines() entities.Identifiers
	HasSuites() bool
	Suites() entities.Identifiers
}
