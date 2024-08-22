package instructions

import (
	"errors"
)

type syscallsBuilder struct {
	list []Syscall
}

func createSyscallsBuilder() SyscallsBuilder {
	out := syscallsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the syscallsBuilder
func (app *syscallsBuilder) Create() SyscallsBuilder {
	return createSyscallsBuilder()
}

// WithList adds a list to the syscallsBuilder
func (app *syscallsBuilder) WithList(list []Syscall) SyscallsBuilder {
	app.list = list
	return app
}

// Now builds a new Syscalls instance
func (app *syscallsBuilder) Now() (Syscalls, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Syscall in order to build a Syscalls instance")
	}

	return createSyscalls(app.list), nil
}
