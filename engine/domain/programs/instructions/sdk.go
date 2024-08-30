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

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// NewReferenceBuilder creates a new reference builder
func NewReferenceBuilder() ReferenceBuilder {
	return createReferenceBuilder()
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
	WithSyscall(syscall Syscall) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Block() string
	Line() uint
	Tokens() Tokens
	HasSyscall() bool
	Syscall() Syscall
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
	WithInstruction(instruction Instruction) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	IsRule() bool
	Rule() rules.Rule
	IsInstruction() bool
	Instruction() Instruction
}

// SyscallBuilder represents the syscall builder
type SyscallBuilder interface {
	Create() SyscallBuilder
	WithFuncName(fnName string) SyscallBuilder
	WithParameters(parameters Parameters) SyscallBuilder
	Now() (Syscall, error)
}

// Syscall represents a syscall
type Syscall interface {
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
	WithName(name string) ParameterBuilder
	WithValue(value Value) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents an execution parameter
type Parameter interface {
	Name() string
	Value() Value
}

// ValueBuilder represents the value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithReference(reference Reference) ValueBuilder
	WithBytes(bytes []byte) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsReference() bool
	Reference() Reference
	IsBytes() bool
	Bytes() []byte
}

// ReferenceBuilder represents the reference builder
type ReferenceBuilder interface {
	Create() ReferenceBuilder
	WithElement(element string) ReferenceBuilder
	WithIndex(index uint) ReferenceBuilder
	Now() (Reference, error)
}

// Reference represents a reference
type Reference interface {
	Element() string
	Index() uint
}
