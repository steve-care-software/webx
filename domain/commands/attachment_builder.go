package commands

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/bytes/criterias"
)

type attachmentBuilder struct {
	global      criterias.Criteria
	local       criterias.Criteria
	application criterias.Criteria
}

func createAttachmentBuilder() AttachmentBuilder {
	out := attachmentBuilder{
		global:      nil,
		local:       nil,
		application: nil,
	}

	return &out
}

// Create initializes the builder
func (app *attachmentBuilder) Create() AttachmentBuilder {
	return createAttachmentBuilder()
}

// WithGlobal adds a global criteria to the builder
func (app *attachmentBuilder) WithGlobal(global criterias.Criteria) AttachmentBuilder {
	app.global = global
	return app
}

// WithLocal adds a local criteria to the builder
func (app *attachmentBuilder) WithLocal(local criterias.Criteria) AttachmentBuilder {
	app.local = local
	return app
}

// WithApplication adds an application criteria to the builder
func (app *attachmentBuilder) WithApplication(application criterias.Criteria) AttachmentBuilder {
	app.application = application
	return app
}

// Now builds a new Attachment instance
func (app *attachmentBuilder) Now() (Attachment, error) {
	if app.global == nil {
		return nil, errors.New("the global criteria is mandatory in order to build an Attachment instance")
	}

	if app.local == nil {
		return nil, errors.New("the local criteria is mandatory in order to build an Attachment instance")
	}

	if app.application == nil {
		return nil, errors.New("the application criteria is mandatory in order to build an Attachment instance")
	}

	return createAttachment(app.global, app.local, app.application), nil
}
