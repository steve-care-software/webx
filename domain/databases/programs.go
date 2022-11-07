package databases

import "github.com/steve-care-software/webx/domain/databases/entities"

type programs struct {
	init   entities.Identifiers
	stop   entities.Identifiers
	start  entities.Identifiers
	daemon entities.Identifiers
}

func createProgramsWithInit(
	init entities.Identifiers,
) Programs {
	return createProgramInternally(init, nil, nil, nil)
}

func createProgramsWithStop(
	stop entities.Identifiers,
) Programs {
	return createProgramInternally(nil, stop, nil, nil)
}

func createProgramsWithStart(
	start entities.Identifiers,
) Programs {
	return createProgramInternally(nil, nil, start, nil)
}

func createProgramsWithDaemon(
	daemon entities.Identifiers,
) Programs {
	return createProgramInternally(nil, nil, nil, daemon)
}

func createProgramsWithInitAndStop(
	init entities.Identifiers,
	stop entities.Identifiers,
) Programs {
	return createProgramInternally(init, stop, nil, nil)
}

func createProgramsWithInitAndStart(
	init entities.Identifiers,
	start entities.Identifiers,
) Programs {
	return createProgramInternally(init, nil, start, nil)
}

func createProgramsWithInitAndDaemon(
	init entities.Identifiers,
	daemon entities.Identifiers,
) Programs {
	return createProgramInternally(init, nil, nil, daemon)
}

func createProgramsWithStopAndStart(
	stop entities.Identifiers,
	start entities.Identifiers,
) Programs {
	return createProgramInternally(nil, stop, start, nil)
}

func createProgramsWithStopAndDaemon(
	stop entities.Identifiers,
	daemon entities.Identifiers,
) Programs {
	return createProgramInternally(nil, stop, nil, daemon)
}

func createProgramsWithStartAndDaemon(
	start entities.Identifiers,
	daemon entities.Identifiers,
) Programs {
	return createProgramInternally(nil, nil, start, daemon)
}

func createProgramsWithInitAndStopAndStart(
	init entities.Identifiers,
	stop entities.Identifiers,
	start entities.Identifiers,
) Programs {
	return createProgramInternally(init, stop, start, nil)
}

func createProgramsWithInitAndStopAndDaemon(
	init entities.Identifiers,
	stop entities.Identifiers,
	daemon entities.Identifiers,
) Programs {
	return createProgramInternally(init, stop, nil, daemon)
}

func createProgramsWithInitAndStartAndDaemon(
	init entities.Identifiers,
	start entities.Identifiers,
	daemon entities.Identifiers,
) Programs {
	return createProgramInternally(init, nil, start, daemon)
}

func createProgramsWithStopAndStartAndDaemon(
	stop entities.Identifiers,
	start entities.Identifiers,
	daemon entities.Identifiers,
) Programs {
	return createProgramInternally(nil, stop, start, daemon)
}

func createProgramsWithInitAndStopAndStartAndDaemon(
	init entities.Identifiers,
	stop entities.Identifiers,
	start entities.Identifiers,
	daemon entities.Identifiers,
) Programs {
	return createProgramInternally(init, stop, start, daemon)
}

func createProgramInternally(
	init entities.Identifiers,
	stop entities.Identifiers,
	start entities.Identifiers,
	daemon entities.Identifiers,
) Programs {
	out := programs{
		init:   init,
		stop:   stop,
		start:  start,
		daemon: daemon,
	}

	return &out
}

// HasInit returns true if there is an init program, false otherwise
func (obj *programs) HasInit() bool {
	return obj.init != nil
}

// Init returns the init, if any
func (obj *programs) Init() entities.Identifiers {
	return obj.init
}

// HasStop returns true if there is a stop program, false otherwise
func (obj *programs) HasStop() bool {
	return obj.stop != nil
}

// Stop returns the stop, if any
func (obj *programs) Stop() entities.Identifiers {
	return obj.stop
}

// HasStart returns true if there is a start program, false otherwise
func (obj *programs) HasStart() bool {
	return obj.start != nil
}

// Start returns the start, if any
func (obj *programs) Start() entities.Identifiers {
	return obj.start
}

// HasDaemon returns true if there is a daemon program, false otherwise
func (obj *programs) HasDaemon() bool {
	return obj.daemon != nil
}

// Daemon returns the daemon, if any
func (obj *programs) Daemon() entities.Identifiers {
	return obj.daemon
}
