package lines

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a line builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithElements(elements []entities.Identifier) Builder
	Now() (Line, error)
}

// Line represents a line
type Line interface {
	Entity() entities.Entity
	Elements() []entities.Identifier
}
