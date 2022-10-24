package programs

import "github.com/steve-care-software/webx/domain/programs/modules"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewApplicationBuilder creates a new application builder instance
func NewApplicationBuilder() ApplicationBuilder {
	return createApplicationBuilder()
}

// NewAttachmentsBuilder creates a new attachments builder
func NewAttachmentsBuilder() AttachmentsBuilder {
	return createAttachmentsBuilder()
}

// NewAttachmentBuilder creates a new attachment builder
func NewAttachmentBuilder() AttachmentBuilder {
	return createAttachmentBuilder()
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	return createAssignmentBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions []Instruction) Builder
	WithOutputs(outputs []string) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Instructions() []Instruction
	HasOutputs() bool
	Outputs() []string
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	WithExecution(execution Application) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsAssignment() bool
	Assignment() Assignment
	IsExecution() bool
	Execution() Application
}

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithName(name string) ApplicationBuilder
	WithModule(module modules.Module) ApplicationBuilder
	WithAttachments(attachments Attachments) ApplicationBuilder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Name() string
	Module() modules.Module
	HasAttachments() bool
	Attachments() Attachments
}

// AttachmentsBuilder represents the attachments builder
type AttachmentsBuilder interface {
	Create() AttachmentsBuilder
	WithList(list []Attachment) AttachmentsBuilder
	Now() (Attachments, error)
}

// Attachments represents attachments
type Attachments interface {
	List() []Attachment
}

// AttachmentBuilder represents an attachment builder
type AttachmentBuilder interface {
	Create() AttachmentBuilder
	WithValue(value Value) AttachmentBuilder
	WithLocal(local string) AttachmentBuilder
	Now() (Attachment, error)
}

// Attachment represents an attachment
type Attachment interface {
	Value() Value
	Local() string
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithIndex(index uint) AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithValue(value Value) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment repesents an assignment
type Assignment interface {
	Index() uint
	Name() string
	Value() Value
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithInput(input string) ValueBuilder
	WithString(str string) ValueBuilder
	WithExecution(execution Application) ValueBuilder
	WithProgram(program Program) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsInput() bool
	Input() string
	IsString() bool
	String() string
	IsExecution() bool
	Execution() Application
	IsProgram() bool
	Program() Program
}
