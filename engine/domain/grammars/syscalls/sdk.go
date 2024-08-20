package syscalls

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/syscalls/values"
)

// Builder represents the syscalls builder
type Builder interface {
	Create() Builder
	WithList(list []Syscall) Builder
	Now() (Syscalls, error)
}

// Syscalls represents syscalls
type Syscalls interface {
	List() []Syscall
}

// SyscallBuilder represents the syscall builder
type SyscallBuilder interface {
	Create() SyscallBuilder
	WithName(name string) SyscallBuilder
	WithFuncName(fnName string) SyscallBuilder
	WithValues(values values.Values) SyscallBuilder
	Now() (Syscall, error)
}

// Syscall represents a syscall
type Syscall interface {
	Name() string
	FuncName() string
	HasValues() bool
	Values() values.Values
}
