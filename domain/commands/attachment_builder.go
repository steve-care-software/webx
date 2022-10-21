package commands

import (
	"errors"

	"github.com/steve-care-software/webx/domain/criterias"
)

type attachmentBuilder struct {
	current     criterias.Criteria
	target      criterias.Criteria
	application criterias.Criteria
}

func createAttachmentBuilder() AttachmentBuilder {
	out := attachmentBuilder{
		current:     nil,
		target:      nil,
		application: nil,
	}

	return &out
}

// Create initializes the builder
func (app *attachmentBuilder) Create() AttachmentBuilder {
	return createAttachmentBuilder()
}

// WithCurrent adds a current criteria to the builder
func (app *attachmentBuilder) WithCurrent(current criterias.Criteria) AttachmentBuilder {
	app.current = current
	return app
}

// WithTarget adds a target criteria to the builder
func (app *attachmentBuilder) WithTarget(target criterias.Criteria) AttachmentBuilder {
	app.target = target
	return app
}

// WithApplication adds an application criteria to the builder
func (app *attachmentBuilder) WithApplication(application criterias.Criteria) AttachmentBuilder {
	app.application = application
	return app
}

// Now builds a new Attachment instance
func (app *attachmentBuilder) Now() (Attachment, error) {
	if app.current == nil {
		return nil, errors.New("the current criteria is mandatory in order to build an Attachment instance")
	}

	if app.target == nil {
		return nil, errors.New("the local criteria is mandatory in order to build an Attachment instance")
	}

	if app.application == nil {
		return nil, errors.New("the application criteria is mandatory in order to build an Attachment instance")
	}

	return createAttachment(app.current, app.target, app.application), nil
}
