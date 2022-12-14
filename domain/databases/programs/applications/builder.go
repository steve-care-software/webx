package applications

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity      entities.Entity
	pModule     *uint
	attachments Attachments
}

func createBuilder() Builder {
	out := builder{
		entity:      nil,
		pModule:     nil,
		attachments: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithEntity adds an entity to the builder
func (app *builder) WithEntity(entity entities.Entity) Builder {
	app.entity = entity
	return app
}

// WithModule adds a module to the builder
func (app *builder) WithModule(module uint) Builder {
	app.pModule = &module
	return app
}

// WithAttachments add attachments to the builder
func (app *builder) WithAttachments(attachments Attachments) Builder {
	app.attachments = attachments
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build an Application instance")
	}

	if app.pModule == nil {
		return nil, errors.New("the module is mandatory in order to build an Application instance")
	}

	if app.attachments != nil {
		return createApplicationWithAttachments(app.entity, *app.pModule, app.attachments), nil
	}

	return createApplication(app.entity, *app.pModule), nil
}
