package instructions

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
	"github.com/steve-care-software/webx/domain/databases/programs/assignments"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithAssignment(assignment assignments.Assignment) Builder
	WithExecution(execution entities.Identifier) Builder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Entity() entities.Entity
	Content() Content
}

// Content represents an instruction's content
type Content interface {
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsExecution() bool
	Execution() entities.Identifier
}
