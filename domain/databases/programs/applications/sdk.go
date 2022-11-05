package applications

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithModule(module entities.Identifier) Builder
	WithAttachments(attachments []Attachment) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Entity() entities.Entity
	Module() entities.Identifier
	HasAttachments() bool
	Attachments() []Attachment
}

// AttachmentBuilder represents an attachment builder
type AttachmentBuilder interface {
	Create() AttachmentBuilder
	WithValue(value entities.Identifier) AttachmentBuilder
	WithLocal(local uint) AttachmentBuilder
	Now() (Attachment, error)
}

// Attachment represents an attachment
type Attachment interface {
	Value() entities.Identifier
	Local() uint
}
