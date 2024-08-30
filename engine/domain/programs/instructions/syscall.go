package instructions

type syscall struct {
	funcName   string
	parameters Parameters
}

func createSyscall(
	funcName string,
) Syscall {
	return createSyscallInternally(funcName, nil)
}

func createSyscallWithParameters(
	funcName string,
	parameters Parameters,
) Syscall {
	return createSyscallInternally(funcName, parameters)
}

func createSyscallInternally(
	funcName string,
	parameters Parameters,
) Syscall {
	out := syscall{
		funcName:   funcName,
		parameters: parameters,
	}

	return &out
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
