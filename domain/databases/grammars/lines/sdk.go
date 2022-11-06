package lines

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a line builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithElements(elements entities.Identifiers) Builder
	Now() (Line, error)
}

// Line represents a line
type Line interface {
	Entity() entities.Entity
	Elements() entities.Identifiers
}
