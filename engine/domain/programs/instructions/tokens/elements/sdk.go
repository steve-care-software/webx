package elements

import "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Builder represents the elements builder
type Builder interface {
	Create() Builder
	WithList(list []Element) Builder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List() []Element
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithRule(rule string) ElementBuilder
	WithSyscall(syscall syscalls.Syscall) ElementBuilder
	WithInstruction(instruction string) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	IsRule() bool
	Rule() string
	IsSyscall() bool
	Syscall() syscalls.Syscall
	IsInstruction() bool
	Instruction() string
}
