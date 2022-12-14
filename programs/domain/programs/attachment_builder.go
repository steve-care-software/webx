package programs

import "errors"

type attachmentBuilder struct {
	value  Value
	pLocal *uint
}

func createAttachmentBuilder() AttachmentBuilder {
	out := attachmentBuilder{
		value:  nil,
		pLocal: nil,
	}

	return &out
}

// Create initializes the builder
func (app *attachmentBuilder) Create() AttachmentBuilder {
	return createAttachmentBuilder()
}

// WithValue adds a value to the builder
func (app *attachmentBuilder) WithValue(value Value) AttachmentBuilder {
	app.value = value
	return app
}

// WithLocal adds a local to the builder
func (app *attachmentBuilder) WithLocal(local uint) AttachmentBuilder {
	app.pLocal = &local
	return app
}

// Now builds a new Attachment instance
func (app *attachmentBuilder) Now() (Attachment, error) {
	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build an Attachment instance")
	}

	if app.pLocal == nil {
		return nil, errors.New("the local is mandatory in order to build an Attachment instance")
	}

	return createAttachment(app.value, *app.pLocal), nil
}
