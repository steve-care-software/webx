package programs

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a program
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithInstructions(instructions entities.Identifiers) Builder
	WithOutputs(outputs []uint) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Entity() entities.Entity
	Instructions() entities.Identifiers
	HasOutputs() bool
	Outputs() []uint
}
