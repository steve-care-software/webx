package programs

import (
	"errors"

	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type applicationBuilder struct {
	name        []byte
	module      modules.Module
	attachments Attachments
}

func createApplicationBuilder() ApplicationBuilder {
	out := applicationBuilder{
		name:        nil,
		module:      nil,
		attachments: nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() ApplicationBuilder {
	return createApplicationBuilder()
}

// WithName adds a name to the builder
func (app *applicationBuilder) WithName(name []byte) ApplicationBuilder {
	app.name = name
	return app
}

// WithModule adds a module to the builder
func (app *applicationBuilder) WithModule(module modules.Module) ApplicationBuilder {
	app.module = module
	return app
}

// WithAttachments add attachments to the builder
func (app *applicationBuilder) WithAttachments(attachments Attachments) ApplicationBuilder {
	app.attachments = attachments
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (Application, error) {
	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	if app.module == nil {
		return nil, errors.New("the module is mandatory in order to build an Application instance")
	}

	if app.attachments != nil {
		return createApplicationWithAttachments(app.name, app.module, app.attachments), nil
	}

	return createApplication(app.name, app.module), nil
}
