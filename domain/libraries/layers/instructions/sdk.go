package instructions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases"
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
	WithAccount(account accounts.Account) InstructionBuilder
	WithDatabase(database databases.Database) InstructionBuilder
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
	IsAccount() bool
	Account() accounts.Account
	IsDatabase() bool
	Database() databases.Database
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
