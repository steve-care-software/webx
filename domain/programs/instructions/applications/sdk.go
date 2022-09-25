package applications

import (
	"github.com/steve-care-software/syntax/domain/programs/instructions/applications/modules"
)

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
	WithAssignment(assignment Assignment) AttachmentBuilder
	WithLocal(local string) AttachmentBuilder
	Now() (Attachment, error)
}

// Attachment represents an attachment
type Attachment interface {
	Assignment() Assignment
	Local() string
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithInput(input string) AssignmentBuilder
	WithValue(value string) AssignmentBuilder
	WithExecution(execution Application) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment repesents an assignment
type Assignment interface {
	Name() string
	Content() AssignmentContent
}

// AssignmentContent represents an assignment content
type AssignmentContent interface {
	IsInput() bool
	Input() string
	IsValue() bool
	Value() string
	IsExecution() bool
	Execution() Application
}
