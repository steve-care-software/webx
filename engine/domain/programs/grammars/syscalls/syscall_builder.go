package syscalls

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
)

type syscallBuilder struct {
	name       string
	funcName   string
	parameters parameters.Parameters
}

func createSyscallBuilder() SyscallBuilder {
	out := syscallBuilder{
		name:       "",
		funcName:   "",
		parameters: nil,
	}

	return &out
}

// Create initializes the builder
func (app *syscallBuilder) Create() SyscallBuilder {
	return createSyscallBuilder()
}

// WithName adds a name to the builder
func (app *syscallBuilder) WithName(name string) SyscallBuilder {
	app.name = name
	return app
}

// WithFuncName adds a name to the builder
func (app *syscallBuilder) WithFuncName(funcName string) SyscallBuilder {
	app.funcName = funcName
	return app
}

// WithParameters add parameters to the builder
func (app *syscallBuilder) WithParameters(parameters parameters.Parameters) SyscallBuilder {
	app.parameters = parameters
	return app
}

// Now builds a new Syscall instance
func (app *syscallBuilder) Now() (Syscall, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Syscall instance")
	}

	if app.funcName == "" {
		return nil, errors.New("the funcName is mandatory in order to build a Syscall instance")
	}

	if app.parameters != nil {
		return createSyscallWithParameters(app.name, app.funcName, app.parameters), nil
	}

	return createSyscall(app.name, app.funcName), nil
}
