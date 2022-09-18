package applications

import (
	"github.com/steve-care-software/logics/domain/programs/instructions/applications/modules"
)

// Application represents an application
type Application interface {
	Name() string
	Module() modules.Module
	HasAttachments() bool
	Attachments() Attachments
}

// Attachments represents attachments
type Attachments interface {
	List() []Attachment
}

// Attachment represents an attachment
type Attachment interface {
	Assignment() Assignment
	Local() string
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
