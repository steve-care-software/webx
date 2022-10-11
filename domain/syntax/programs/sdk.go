package programs

import (
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithAssignments(assignments []applications.Assignment) Builder
	WithOutputs(outputs []string) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Assignments() []applications.Assignment
	HasOutputs() bool
	Outputs() []string
}
