package instructions

import (
	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions/applications"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// Builder represents an instructions builder
type Builder interface {
	Create() Builder
	WithList(list []Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithAssignment(assignment applications.Assignment) InstructionBuilder
	WithExecution(execution applications.Application) InstructionBuilder
	WithDelete(delete applications.Application) InstructionBuilder
	WithSetPath(setPath string) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsAssignment() bool
	Assignment() applications.Assignment
	IsExecution() bool
	Execution() applications.Application
	IsDelete() bool
	Delete() applications.Application
	IsSetPath() bool
	SetPath() string
}
