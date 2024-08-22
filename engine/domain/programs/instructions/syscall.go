package instructions

type syscall struct {
	name       string
	funcName   string
	parameters Parameters
}

func createSyscall(
	name string,
	funcName string,
) Syscall {
	return createSyscallInternally(name, funcName, nil)
}

func createSyscallWithParameters(
	name string,
	funcName string,
	parameters Parameters,
) Syscall {
	return createSyscallInternally(name, funcName, parameters)
}

func createSyscallInternally(
	name string,
	funcName string,
	parameters Parameters,
) Syscall {
	out := syscall{
		name:       name,
		funcName:   funcName,
		parameters: parameters,
	}

	return &out
}

// Name returns the name
func (obj *syscall) Name() string {
	return obj.name
}

// FuncName returns the funcName
func (obj *syscall) FuncName() string {
	return obj.funcName
}

// HasParameters returns true if there is parameters, false otherwise
func (obj *syscall) HasParameters() bool {
	return obj.parameters != nil
}

// Parameters returns the parameters
func (obj *syscall) Parameters() Parameters {
	return obj.parameters
}
