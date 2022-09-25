package applications

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/programs/instructions/applications/modules"
)

type builder struct {
	name        string
	module      modules.Module
	attachments Attachments
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		module:      nil,
		attachments: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithModule adds a module to the builder
func (app *builder) WithModule(module modules.Module) Builder {
	app.module = module
	return app
}

// WithAttachments add attachments to the builder
func (app *builder) WithAttachments(attachments Attachments) Builder {
	app.attachments = attachments
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.name == "" {
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
