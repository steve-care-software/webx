package syscalls

import (
	"errors"
	"fmt"
)

type syscalls struct {
	list []Syscall
	mp   map[string]Syscall
}

func createSyscalls(
	list []Syscall,
	mp map[string]Syscall,
) Syscalls {
	out := syscalls{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of syscall
func (obj *syscalls) List() []Syscall {
	return obj.list
}

// Fetch fetches a syscall by name
func (obj *syscalls) Fetch(name string) (Syscall, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the syscall (name: %s) does not exists", name)
	return nil, errors.New(str)
}
