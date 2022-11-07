package applications

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
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

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithModule(module entities.Identifier) Builder
	WithAttachments(attachments Attachments) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Entity() entities.Entity
	Module() entities.Identifier
	HasAttachments() bool
	Attachments() Attachments
}

// AttachmentsBuilder represents an attachments builder
type AttachmentsBuilder interface {
	Create() AttachmentsBuilder
	WithList(list []Attachment) AttachmentsBuilder
	Now() (Attachments, error)
}

// Attachments represents the attachments
type Attachments interface {
	List() []Attachment
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
