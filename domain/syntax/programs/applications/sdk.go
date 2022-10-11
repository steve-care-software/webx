package applications

import (
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications/modules"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithModule(module modules.Module) Builder
	WithAttachments(attachments Attachments) Builder
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
}
