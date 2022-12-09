package applications

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type attachmentBuilder struct {
	pValue *hash.Hash
	pLocal *uint
}

func createAttachmentBuilder() AttachmentBuilder {
	out := attachmentBuilder{
		pValue: nil,
		pLocal: nil,
	}

	return &out
}

// Create initializes the builder
func (app *attachmentBuilder) Create() AttachmentBuilder {
	return createAttachmentBuilder()
}

// WithValue adds a value to the builder
func (app *attachmentBuilder) WithValue(value hash.Hash) AttachmentBuilder {
	app.pValue = &value
	return app
}

// WithLocal adds a local to the builder
func (app *attachmentBuilder) WithLocal(local uint) AttachmentBuilder {
	app.pLocal = &local
	return app
}

// Now builds a new Attachment instance
func (app *attachmentBuilder) Now() (Attachment, error) {
	if app.pValue == nil {
		return nil, errors.New("the value is mandatory in order to build an Attachment instance")
	}

	if app.pLocal == nil {
		return nil, errors.New("the local is mandatory in order to build an Attachment instance")
	}

	return createAttachment(*app.pValue, *app.pLocal), nil
}
