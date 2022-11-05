package values

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
	"github.com/steve-care-software/webx/domain/databases/programs/assignments"
)

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithInput(input uint) Builder
	WithAssignment(assignment assignments.Assignment) Builder
	WithExecution(execution entities.Identifier) Builder
	WithProgram(program entities.Identifier) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Entity() entities.Entity
	Content() Content
}

// Content represents a value content
type Content interface {
	IsInput() bool
	Input() uint
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsExecution() bool
	Execution() entities.Identifier
	IsProgram() bool
	Program() entities.Identifier
}
