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

// NewOutputBuilder creates a new output builder
func NewOutputBuilder() OutputBuilder {
	return createOutputBuilder()
}

// Builder represents instructions builder
type Builder interface {
	Create() Builder
	WithList(instructions []Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithModule(module string) InstructionBuilder
	WithApplication(application applications.Application) InstructionBuilder
	WithParameter(parameter parameters.Parameter) InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	WithExecution(execution string) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsModule() bool
	Module() string
	IsApplication() bool
	Application() applications.Application
	IsParameter() bool
	Parameter() parameters.Parameter
	IsAssignment() bool
	Assignment() Assignment
	IsAttachment() bool
	Attachment() attachments.Attachment
	IsExecution() bool
	Execution() string
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithVariable(variable string) AssignmentBuilder
	WithValue(value Value) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Variable() string
	Value() Value
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithInput(input string) ValueBuilder
	WithString(str string) ValueBuilder
	WithInstructions(instructions Instructions) ValueBuilder
	WithExecution(execution string) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsInput() bool
	Input() string
	IsString() bool
	String() string
	IsInstructions() bool
	Instructions() Instructions
	IsExecution() bool
	Execution() string
}

// OutputBuilder represents an output builder
type OutputBuilder interface {
	Create() OutputBuilder
	WithInstructions(instructions Instructions) OutputBuilder
	WithRemaining(remaining []byte) OutputBuilder
	Now() (Output, error)
}

// Output represents an output
type Output interface {
	Instructions() Instructions
	HasRemaining() bool
	Remaining() []byte
}
