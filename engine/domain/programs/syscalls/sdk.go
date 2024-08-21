package syscalls

import "github.com/steve-care-software/webx/engine/domain/programs/syscalls/values"

// Syscalls represents syscalls
type Syscalls interface {
	List() []Syscall
}

// Syscall represents a syscall
type Syscall interface {
	Name() string
	FuncName() string
	HasValues() bool
	Values() values.Values
}
