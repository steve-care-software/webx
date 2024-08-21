package instructions

import "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
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
	Fetch(name string) (Instruction, error)
}

// InstructionBuilder represents the instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithBlock(block string) InstructionBuilder
	WithLine(line uint) InstructionBuilder
	WithTokens(tokens tokens.Tokens) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Block() string
	Line() uint
	Tokens() tokens.Tokens
}
