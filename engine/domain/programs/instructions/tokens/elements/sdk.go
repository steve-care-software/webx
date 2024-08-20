package elements

// Builder represents the element builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithRule(rule string) Builder
	WithSyscall(syscall string) Builder
	WithInstruction(instruction string) Builder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	IsRule() bool
	Rule() string
	IsSyscall() bool
	Syscall() string
	IsInstruction() bool
	Instruction() string
}
