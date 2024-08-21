package syscalls

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewSyscallBuilder creates a new syscall builder
func NewSyscallBuilder() SyscallBuilder {
	return createSyscallBuilder()
}

// Builder represents the syscalls builder
type Builder interface {
	Create() Builder
	WithList(list []Syscall) Builder
	Now() (Syscalls, error)
}

// Syscalls represents syscalls
type Syscalls interface {
	List() []Syscall
	Fetch(name string) (Syscall, error)
}

// SyscallBuilder represents the syscall builder
type SyscallBuilder interface {
	Create() SyscallBuilder
	WithName(name string) SyscallBuilder
	WithFuncName(fnName string) SyscallBuilder
	WithParameters(parameters parameters.Parameters) SyscallBuilder
	Now() (Syscall, error)
}

// Syscall represents a syscall
type Syscall interface {
	Name() string
	FuncName() string
	HasParameters() bool
	Parameters() parameters.Parameters
}
