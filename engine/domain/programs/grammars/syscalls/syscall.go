package syscalls

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls/values"

type syscall struct {
	name     string
	funcName string
	values   values.Values
}

func createSyscall(
	name string,
	funcName string,
) Syscall {
	return createSyscallInternally(name, funcName, nil)
}

func createSyscallWithValues(
	name string,
	funcName string,
	values values.Values,
) Syscall {
	return createSyscallInternally(name, funcName, values)
}

func createSyscallInternally(
	name string,
	funcName string,
	values values.Values,
) Syscall {
	out := syscall{
		name:     name,
		funcName: funcName,
		values:   values,
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

// HasValues returns true if there is values, false otherwise
func (obj *syscall) HasValues() bool {
	return obj.values != nil
}

// Values returns the values
func (obj *syscall) Values() values.Values {
	return obj.values
}
