package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
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
	WithHash(hash hash.Hash) Builder
	WithModule(module uint) Builder
	WithAttachments(attachments Attachments) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Hash() hash.Hash
	Module() uint
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
	WithValue(value hash.Hash) AttachmentBuilder
	WithLocal(local uint) AttachmentBuilder
	Now() (Attachment, error)
}

// Attachment represents an attachment
type Attachment interface {
	Value() hash.Hash
	Local() uint
}
