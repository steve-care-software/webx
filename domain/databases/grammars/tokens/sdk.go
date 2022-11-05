package tokens

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a token builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithLines(lines []entities.Identifier) Builder
	WithSuites(suites []entities.Identifier) Builder
	Now() (Token, error)
}

// Token represents token metadata
type Token interface {
	Entity() entities.Entity
	Lines() []entities.Identifier
	HasSuites() bool
	Suites() []entities.Identifier
}
