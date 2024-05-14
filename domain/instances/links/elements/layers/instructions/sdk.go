package instructions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments"
)

// NewBuilder creates a new instructions builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	hashAdapter := hash.NewAdapter()
	return createInstructionBuilder(
		hashAdapter,
	)
}

// NewConditionBuilder creates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionBuilder(
		hashAdapter,
	)
}

// NewLoopBuuilder creates a new loop builder
func NewLoopBuuilder() LoopBuilder {
	hashAdapter := hash.NewAdapter()
	return createLoopBuilder(
		hashAdapter,
	)
}

// Adapter represents the instructions adapter
type Adapter interface {
	ToBytes(ins Instructions) ([]byte, error)
	ToInstance(bytes []byte) (Instructions, error)
}

// Builder represents instructions builder
type Builder interface {
	Create() Builder
	WithList(list []Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	Hash() hash.Hash
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithRaiseError(raiseError uint) InstructionBuilder
	WithCondition(condition Condition) InstructionBuilder
	WithAssignment(assignment assignments.Assignment) InstructionBuilder
	IsStop() InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Hash() hash.Hash
	IsStop() bool
	IsRaiseError() bool
	RaiseError() uint
	IsCondition() bool
	Condition() Condition
	IsAssignment() bool
	Assignment() assignments.Assignment
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithVariable(variable string) ConditionBuilder
	WithInstructions(instructions Instructions) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Variable() string
	Instructions() Instructions
}

// LoopBuilder represents the loop builder
type LoopBuilder interface {
	Create() LoopBuilder
	WithAmount(amount string) LoopBuilder
	WithInstructions(instructions Instructions) LoopBuilder
	Now() (Loop, error)
}

// Loop represents a loop
type Loop interface {
	Hash() hash.Hash
	Amount() string
	Instructions() Instructions
}
