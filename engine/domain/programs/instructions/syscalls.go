package instructions

type syscalls struct {
	list []Syscall
}

func createSyscalls(
	list []Syscall,
) Syscalls {
	out := syscalls{
		list: list,
	}

	return &out
}

// List returns the list of syscall
func (obj *syscalls) List() []Syscall {
	return obj.list
}
