package applications

import "errors"

type attachmentBuilder struct {
	assignment Assignment
	local      string
}

func createAttachmentBuilder() AttachmentBuilder {
	out := attachmentBuilder{
		assignment: nil,
		local:      "",
	}

	return &out
}

// Create initializes the builder
func (app *attachmentBuilder) Create() AttachmentBuilder {
	return createAttachmentBuilder()
}

// WithAssignment adds an assignment to the builder
func (app *attachmentBuilder) WithAssignment(assignment Assignment) AttachmentBuilder {
	app.assignment = assignment
	return app
}

// WithLocal adds a local to the builder
func (app *attachmentBuilder) WithLocal(local string) AttachmentBuilder {
	app.local = local
	return app
}

// Now builds a new Attachment instance
func (app *attachmentBuilder) Now() (Attachment, error) {
	if app.assignment == nil {
		return nil, errors.New("the assignment is mandatory in order to build an Attachment instance")
	}

	if app.local == "" {
		return nil, errors.New("the local is mandatory in order to build an Attachment instance")
	}

	return createAttachment(app.assignment, app.local), nil
}
