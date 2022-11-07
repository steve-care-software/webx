package databases

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type programsBuilder struct {
	init   entities.Identifiers
	stop   entities.Identifiers
	start  entities.Identifiers
	daemon entities.Identifiers
}

func createProgramsBuilder() ProgramsBuilder {
	out := programsBuilder{
		init:   nil,
		stop:   nil,
		start:  nil,
		daemon: nil,
	}

	return &out
}

// Create initializes the builder
func (app *programsBuilder) Create() ProgramsBuilder {
	return createProgramsBuilder()
}

// WithInit adds an init to the builder
func (app *programsBuilder) WithInit(init entities.Identifiers) ProgramsBuilder {
	app.init = init
	return app
}

// WithStop adds a stop to the builder
func (app *programsBuilder) WithStop(stop entities.Identifiers) ProgramsBuilder {
	app.stop = stop
	return app
}

// WithStart adds a start to the builder
func (app *programsBuilder) WithStart(start entities.Identifiers) ProgramsBuilder {
	app.start = start
	return app
}

// WithDaemon adds a daemon to the builder
func (app *programsBuilder) WithDaemon(daemon entities.Identifiers) ProgramsBuilder {
	app.daemon = daemon
	return app
}

// Now builds a new Programs instance
func (app *programsBuilder) Now() (Programs, error) {
	if app.init != nil && app.stop != nil && app.start != nil && app.daemon != nil {
		return createProgramsWithInitAndStopAndStartAndDaemon(app.init, app.stop, app.start, app.daemon), nil
	}

	if app.init != nil && app.start != nil && app.daemon != nil {
		return createProgramsWithInitAndStartAndDaemon(app.init, app.start, app.daemon), nil
	}

	if app.init != nil && app.stop != nil && app.daemon != nil {
		return createProgramsWithInitAndStopAndDaemon(app.init, app.stop, app.daemon), nil
	}

	if app.init != nil && app.stop != nil && app.start != nil {
		return createProgramsWithInitAndStopAndStart(app.init, app.stop, app.start), nil
	}

	if app.stop != nil && app.start != nil && app.daemon != nil {
		return createProgramsWithStopAndStartAndDaemon(app.stop, app.start, app.daemon), nil
	}

	if app.init != nil && app.daemon != nil {
		return createProgramsWithInitAndDaemon(app.init, app.daemon), nil
	}

	if app.init != nil && app.stop != nil {
		return createProgramsWithInitAndStop(app.init, app.stop), nil
	}

	if app.init != nil && app.start != nil {
		return createProgramsWithInitAndStart(app.init, app.start), nil
	}

	if app.stop != nil && app.start != nil {
		return createProgramsWithStopAndStart(app.stop, app.start), nil
	}

	if app.stop != nil && app.daemon != nil {
		return createProgramsWithStopAndDaemon(app.stop, app.daemon), nil
	}

	if app.start != nil && app.daemon != nil {
		return createProgramsWithStartAndDaemon(app.start, app.daemon), nil
	}

	if app.init != nil {
		return createProgramsWithInit(app.init), nil
	}

	if app.start != nil {
		return createProgramsWithStart(app.start), nil
	}

	if app.stop != nil {
		return createProgramsWithStop(app.stop), nil
	}

	if app.daemon != nil {
		return createProgramsWithDaemon(app.daemon), nil
	}

	return nil, errors.New("the Programs is invalid")
}

/*

init   entities.Identifiers
stop   entities.Identifiers
start  entities.Identifiers
daemon entities.Identifiers

*/
