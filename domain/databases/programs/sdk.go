package programs

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a program
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithInstructions(instructions []entities.Identifier) Builder
	WithOutputs(outputs []uint) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Entity() entities.Entity
	Instructions() []entities.Identifier
	HasOutputs() bool
	Outputs() []uint
}
