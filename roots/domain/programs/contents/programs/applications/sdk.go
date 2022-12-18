package applications

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
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

// Adapter represents the application adapter
type Adapter interface {
	ToContent(ins Application) ([]byte, error)
	ToApplication(content []byte) (Application, error)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithIndex(index uint) Builder
	WithModule(module uint) Builder
	WithAttachments(attachments Attachments) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Hash() hash.Hash
	Index() uint
	Module() uint
	HasAttachments() bool
	Attachments() Attachments
}

// AttachmentsAdapter represents the attachments adapter
type AttachmentsAdapter interface {
	ToContent(ins Attachments) ([]byte, error)
	ToAttachments(content []byte) (Attachments, error)
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

// AttachmentAdapter represents the attachment adapter
type AttachmentAdapter interface {
	ToContent(ins Attachment) ([]byte, error)
	ToAttachment(content []byte) (Attachment, error)
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
