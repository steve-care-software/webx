package instructions

import (
	"errors"
)

type syscallBuilder struct {
	funcName   string
	parameters Parameters
}

func createSyscallBuilder() SyscallBuilder {
	out := syscallBuilder{
		funcName:   "",
		parameters: nil,
	}

	return &out
}

// Create initializes the builder
func (app *syscallBuilder) Create() SyscallBuilder {
	return createSyscallBuilder()
}

// WithFuncName adds a name to the builder
func (app *syscallBuilder) WithFuncName(funcName string) SyscallBuilder {
	app.funcName = funcName
	return app
}

// WithParameters add parameters to the builder
func (app *syscallBuilder) WithParameters(parameters Parameters) SyscallBuilder {
	app.parameters = parameters
	return app
}

// Now builds a new Syscall instance
func (app *syscallBuilder) Now() (Syscall, error) {
	if app.funcName == "" {
		return nil, errors.New("the funcName is mandatory in order to build a Syscall instance")
	}

	if app.parameters != nil {
		return createSyscallWithParameters(app.funcName, app.parameters), nil
	}

	return createSyscall(app.funcName), nil
}
