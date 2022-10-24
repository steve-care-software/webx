package programs

import "errors"

type attachmentsBuilder struct {
	list []Attachment
}

func createAttachmentsBuilder() AttachmentsBuilder {
	out := attachmentsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *attachmentsBuilder) Create() AttachmentsBuilder {
	return createAttachmentsBuilder()
}

// WithList adds a list to the builder
func (app *attachmentsBuilder) WithList(list []Attachment) AttachmentsBuilder {
	app.list = list
	return app
}

// Now builds a new Attachments instance
func (app *attachmentsBuilder) Now() (Attachments, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Attachment in order to build a Attachments instance")
	}

	return createAttachments(app.list), nil
}
