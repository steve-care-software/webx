package applications

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity entities.Entity
	init   entities.Identifiers
	stop   entities.Identifiers
	start  entities.Identifiers
	daemon entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		entity: nil,
		init:   nil,
		stop:   nil,
		start:  nil,
		daemon: nil,
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

// WithInit adds an init to the builder
func (app *builder) WithInit(init entities.Identifiers) Builder {
	app.init = init
	return app
}

// WithStop adds a stop to the builder
func (app *builder) WithStop(stop entities.Identifiers) Builder {
	app.stop = stop
	return app
}

// WithStart adds a start to the builder
func (app *builder) WithStart(start entities.Identifiers) Builder {
	app.start = start
	return app
}

// WithDaemon adds a daemon to the builder
func (app *builder) WithDaemon(daemon entities.Identifiers) Builder {
	app.daemon = daemon
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build an Application instance")
	}

	if app.init != nil && app.stop != nil && app.start != nil && app.daemon != nil {
		content := createContentWithInitAndStopAndStartAndDaemon(app.init, app.stop, app.start, app.daemon)
		return createApplication(app.entity, content), nil
	}

	if app.init != nil && app.start != nil && app.daemon != nil {
		content := createContentWithInitAndStartAndDaemon(app.init, app.start, app.daemon)
		return createApplication(app.entity, content), nil
	}

	if app.init != nil && app.stop != nil && app.daemon != nil {
		content := createContentWithInitAndStopAndDaemon(app.init, app.stop, app.daemon)
		return createApplication(app.entity, content), nil
	}

	if app.init != nil && app.stop != nil && app.start != nil {
		content := createContentWithInitAndStopAndStart(app.init, app.stop, app.start)
		return createApplication(app.entity, content), nil
	}

	if app.stop != nil && app.start != nil && app.daemon != nil {
		content := createContentWithStopAndStartAndDaemon(app.stop, app.start, app.daemon)
		return createApplication(app.entity, content), nil
	}

	if app.init != nil && app.daemon != nil {
		content := createContentWithInitAndDaemon(app.init, app.daemon)
		return createApplication(app.entity, content), nil
	}

	if app.init != nil && app.stop != nil {
		content := createContentWithInitAndStop(app.init, app.stop)
		return createApplication(app.entity, content), nil
	}

	if app.init != nil && app.start != nil {
		content := createContentWithInitAndStart(app.init, app.start)
		return createApplication(app.entity, content), nil
	}

	if app.stop != nil && app.start != nil {
		content := createContentWithStopAndStart(app.stop, app.start)
		return createApplication(app.entity, content), nil
	}

	if app.stop != nil && app.daemon != nil {
		content := createContentWithStopAndDaemon(app.stop, app.daemon)
		return createApplication(app.entity, content), nil
	}

	if app.start != nil && app.daemon != nil {
		content := createContentWithStartAndDaemon(app.start, app.daemon)
		return createApplication(app.entity, content), nil
	}

	if app.init != nil {
		content := createContentWithInit(app.init)
		return createApplication(app.entity, content), nil
	}

	if app.start != nil {
		content := createContentWithStart(app.start)
		return createApplication(app.entity, content), nil
	}

	if app.stop != nil {
		content := createContentWithStop(app.stop)
		return createApplication(app.entity, content), nil
	}

	if app.daemon != nil {
		content := createContentWithDaemon(app.daemon)
		return createApplication(app.entity, content), nil
	}

	return nil, errors.New("the Application is invalid")
}
