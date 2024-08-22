package instructions

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewTokensBuilder creates a new tokens builder
func NewTokensBuilder() TokensBuilder {
	return createTokensBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewElementsAdapter creates a new elements adapter
func NewElementsAdapter() ElementsAdapter {
	return createElementsAdapter()
}

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	return createElementsBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewSyscallsBuilder creates a new syscalls builder
func NewSyscallsBuilder() SyscallsBuilder {
	return createSyscallsBuilder()
}

// NewSyscallBuilder creates a new syscall builder
func NewSyscallBuilder() SyscallBuilder {
	return createSyscallBuilder()
}

// NewParametersBuilder creates a new parameters builder
func NewParametersBuilder() ParametersBuilder {
	return createParametersBuilder()
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	return createParameterBuilder()
}

// Builder represents the instructions builder
type Builder interface {
	Create() Builder
	WithList(list []Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
	Fetch(name string, idx uint) (Instruction, error)
}

// InstructionBuilder represents the instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithBlock(block string) InstructionBuilder
	WithLine(line uint) InstructionBuilder
	WithTokens(tokens Tokens) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Block() string
	Line() uint
	Tokens() Tokens
}

// TokensBuilder represents the tokens builder
type TokensBuilder interface {
	Create() TokensBuilder
	WithList(list []Token) TokensBuilder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	List() []Token
	Fetch(name string, index uint) (Token, error)
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithElements(elements Elements) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Elements() Elements
}

// ElementsAdapter represents the elements adapter
type ElementsAdapter interface {
	// ToBytes takes an elements and returns its bytes
	ToBytes(elements Elements) ([]byte, error)
}

// ElementsBuilder represents the elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List() []Element
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithRule(rule rules.Rule) ElementBuilder
	WithSyscall(syscall Syscall) ElementBuilder
	WithInstruction(instruction Instruction) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	IsRule() bool
	Rule() rules.Rule
	IsSyscall() bool
	Syscall() Syscall
	IsInstruction() bool
	Instruction() Instruction
}

// SyscallsBuilder represents the syscalls builder
type SyscallsBuilder interface {
	Create() SyscallsBuilder
	WithList(list []Syscall) SyscallsBuilder
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
	WithParameters(parameters Parameters) SyscallBuilder
	Now() (Syscall, error)
}

// Syscall represents a syscall
type Syscall interface {
	Name() string
	FuncName() string
	HasParameters() bool
	Parameters() Parameters
}

// ParametersBuilder represents the parameters builder
type ParametersBuilder interface {
	Create() ParametersBuilder
	WithList(list []Parameter) ParametersBuilder
	Now() (Parameters, error)
}

// Parameters represents parameters
type Parameters interface {
	List() []Parameter
}

// ParameterBuilder represents the parameter builder
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithElement(element string) ParameterBuilder
	WithIndex(index uint) ParameterBuilder
	WithName(name string) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents an execution parameter
type Parameter interface {
	Element() string
	Index() uint
	Name() string
}
