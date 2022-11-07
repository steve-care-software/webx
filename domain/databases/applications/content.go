package applications

import "github.com/steve-care-software/webx/domain/databases/entities"

type content struct {
	init   entities.Identifiers
	stop   entities.Identifiers
	start  entities.Identifiers
	daemon entities.Identifiers
}

func createContentWithInit(
	init entities.Identifiers,
) Content {
	return createProgramInternally(init, nil, nil, nil)
}

func createContentWithStop(
	stop entities.Identifiers,
) Content {
	return createProgramInternally(nil, stop, nil, nil)
}

func createContentWithStart(
	start entities.Identifiers,
) Content {
	return createProgramInternally(nil, nil, start, nil)
}

func createContentWithDaemon(
	daemon entities.Identifiers,
) Content {
	return createProgramInternally(nil, nil, nil, daemon)
}

func createContentWithInitAndStop(
	init entities.Identifiers,
	stop entities.Identifiers,
) Content {
	return createProgramInternally(init, stop, nil, nil)
}

func createContentWithInitAndStart(
	init entities.Identifiers,
	start entities.Identifiers,
) Content {
	return createProgramInternally(init, nil, start, nil)
}

func createContentWithInitAndDaemon(
	init entities.Identifiers,
	daemon entities.Identifiers,
) Content {
	return createProgramInternally(init, nil, nil, daemon)
}

func createContentWithStopAndStart(
	stop entities.Identifiers,
	start entities.Identifiers,
) Content {
	return createProgramInternally(nil, stop, start, nil)
}

func createContentWithStopAndDaemon(
	stop entities.Identifiers,
	daemon entities.Identifiers,
) Content {
	return createProgramInternally(nil, stop, nil, daemon)
}

func createContentWithStartAndDaemon(
	start entities.Identifiers,
	daemon entities.Identifiers,
) Content {
	return createProgramInternally(nil, nil, start, daemon)
}

func createContentWithInitAndStopAndStart(
	init entities.Identifiers,
	stop entities.Identifiers,
	start entities.Identifiers,
) Content {
	return createProgramInternally(init, stop, start, nil)
}

func createContentWithInitAndStopAndDaemon(
	init entities.Identifiers,
	stop entities.Identifiers,
	daemon entities.Identifiers,
) Content {
	return createProgramInternally(init, stop, nil, daemon)
}

func createContentWithInitAndStartAndDaemon(
	init entities.Identifiers,
	start entities.Identifiers,
	daemon entities.Identifiers,
) Content {
	return createProgramInternally(init, nil, start, daemon)
}

func createContentWithStopAndStartAndDaemon(
	stop entities.Identifiers,
	start entities.Identifiers,
	daemon entities.Identifiers,
) Content {
	return createProgramInternally(nil, stop, start, daemon)
}

func createContentWithInitAndStopAndStartAndDaemon(
	init entities.Identifiers,
	stop entities.Identifiers,
	start entities.Identifiers,
	daemon entities.Identifiers,
) Content {
	return createProgramInternally(init, stop, start, daemon)
}

func createProgramInternally(
	init entities.Identifiers,
	stop entities.Identifiers,
	start entities.Identifiers,
	daemon entities.Identifiers,
) Content {
	out := content{
		init:   init,
		stop:   stop,
		start:  start,
		daemon: daemon,
	}

	return &out
}

// HasInit returns true if there is an init program, false otherwise
func (obj *content) HasInit() bool {
	return obj.init != nil
}

// Init returns the init, if any
func (obj *content) Init() entities.Identifiers {
	return obj.init
}

// HasStop returns true if there is a stop program, false otherwise
func (obj *content) HasStop() bool {
	return obj.stop != nil
}

// Stop returns the stop, if any
func (obj *content) Stop() entities.Identifiers {
	return obj.stop
}

// HasStart returns true if there is a start program, false otherwise
func (obj *content) HasStart() bool {
	return obj.start != nil
}

// Start returns the start, if any
func (obj *content) Start() entities.Identifiers {
	return obj.start
}

// HasDaemon returns true if there is a daemon program, false otherwise
func (obj *content) HasDaemon() bool {
	return obj.daemon != nil
}

// Daemon returns the daemon, if any
func (obj *content) Daemon() entities.Identifiers {
	return obj.daemon
}
