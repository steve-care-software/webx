package instructions

import (
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	return createAssignmentBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// Builder represents instructions builder
type Builder interface {
	Create() Builder
	WithList(instructions []Instruction) Builder
	WithRemaining(remaining []byte) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
	HasRemaining() bool
	Remaining() []byte
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithModule(module []byte) InstructionBuilder
	WithApplication(application applications.Application) InstructionBuilder
	WithParameter(parameter parameters.Parameter) InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	WithAttachment(attachment attachments.Attachment) InstructionBuilder
	WithExecution(execution []byte) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsModule() bool
	Module() []byte
	IsApplication() bool
	Application() applications.Application
	IsParameter() bool
	Parameter() parameters.Parameter
	IsAssignment() bool
	Assignment() Assignment
	IsAttachment() bool
	Attachment() attachments.Attachment
	IsExecution() bool
	Execution() []byte
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithVariable(variable []byte) AssignmentBuilder
	WithValue(value Value) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Variable() []byte
	Value() Value
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithVariable(variable []byte) ValueBuilder
	WithConstant(constant []byte) ValueBuilder
	WithInstructions(instructions Instructions) ValueBuilder
	WithExecution(execution []byte) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsVariable() bool
	Variable() []byte
	IsConstant() bool
	Constant() []byte
	IsInstructions() bool
	Instructions() Instructions
	IsExecution() bool
	Execution() []byte
}
