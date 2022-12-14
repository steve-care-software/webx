package programs

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewInstructionsBuilder creates a new instructions builder
func NewInstructionsBuilder() InstructionsBuilder {
	hashAdapter := hash.NewAdapter()
	return createInstructionsBuilder(hashAdapter)
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewApplicationBuilder creates a new application builder instance
func NewApplicationBuilder() ApplicationBuilder {
	hashAdapter := hash.NewAdapter()
	return createApplicationBuilder(
		hashAdapter,
	)
}

// NewAttachmentsBuilder creates a new attachments builder
func NewAttachmentsBuilder() AttachmentsBuilder {
	hashAdapter := hash.NewAdapter()
	return createAttachmentsBuilder(
		hashAdapter,
	)
}

// NewAttachmentBuilder creates a new attachment builder
func NewAttachmentBuilder() AttachmentBuilder {
	return createAttachmentBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createValueBuilder(
		hashAdapter,
	)
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions Instructions) Builder
	WithOutputs(outputs []uint) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Instructions() Instructions
	HasOutputs() bool
	Outputs() []uint
}

// InstructionsBuilder represents instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithList(list []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	Hash() hash.Hash
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithValue(value Value) InstructionBuilder
	WithExecution(execution Application) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Hash() hash.Hash
	IsValue() bool
	Value() Value
	IsExecution() bool
	Execution() Application
}

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithIndex(index uint) ApplicationBuilder
	WithModule(module modules.Module) ApplicationBuilder
	WithAttachments(attachments Attachments) ApplicationBuilder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Hash() hash.Hash
	Index() uint
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
	Hash() hash.Hash
	List() []Attachment
}

// AttachmentBuilder represents an attachment builder
type AttachmentBuilder interface {
	Create() AttachmentBuilder
	WithValue(value Value) AttachmentBuilder
	WithLocal(local uint) AttachmentBuilder
	Now() (Attachment, error)
}

// Attachment represents an attachment
type Attachment interface {
	Value() Value
	Local() uint
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithInput(input uint) ValueBuilder
	WithConstant(constant []byte) ValueBuilder
	WithExecution(execution Application) ValueBuilder
	WithProgram(program Program) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Hash() hash.Hash
	Content() Content
}

// Content represents a value's content
type Content interface {
	IsInput() bool
	Input() *uint
	IsConstant() bool
	Constant() []byte
	IsExecution() bool
	Execution() Application
	IsProgram() bool
	Program() Program
}
