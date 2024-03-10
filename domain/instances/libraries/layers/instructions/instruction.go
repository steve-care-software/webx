package instructions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases"
)

type instruction struct {
	hash       hash.Hash
	isStop     bool
	raiseError uint
	condition  Condition
	assignment assignments.Assignment
	account    accounts.Account
	database   databases.Database
}

func createInstructionWithIsStop(
	hash hash.Hash,
) Instruction {
	return createInstructionInternally(
		hash,
		true,
		0,
		nil,
		nil,
		nil,
		nil,
	)
}

func createInstructionWithRaiseError(
	hash hash.Hash,
	raiseError uint,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		raiseError,
		nil,
		nil,
		nil,
		nil,
	)
}

func createInstructionWithCondition(
	hash hash.Hash,
	condition Condition,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		condition,
		nil,
		nil,
		nil,
	)
}

func createInstructionWithAssignment(
	hash hash.Hash,
	assignment assignments.Assignment,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		assignment,
		nil,
		nil,
	)
}

func createInstructionWithAccount(
	hash hash.Hash,
	account accounts.Account,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		nil,
		account,
		nil,
	)
}

func createInstructionWithDatabase(
	hash hash.Hash,
	database databases.Database,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		nil,
		nil,
		database,
	)
}

func createInstructionInternally(
	hash hash.Hash,
	isStop bool,
	raiseError uint,
	condition Condition,
	assignment assignments.Assignment,
	account accounts.Account,
	database databases.Database,
) Instruction {
	out := instruction{
		hash:       hash,
		isStop:     isStop,
		raiseError: raiseError,
		condition:  condition,
		assignment: assignment,
		account:    account,
		database:   database,
	}

	return &out
}

// Hash returns the hash
func (obj *instruction) Hash() hash.Hash {
	return obj.hash
}

// IsStop returns true if stop, false otherwise
func (obj *instruction) IsStop() bool {
	return obj.isStop
}

// IsRaiseError returns true if raiseError, false otherwise
func (obj *instruction) IsRaiseError() bool {
	return obj.raiseError > 0
}

// RaiseError returns the raiseError, if any
func (obj *instruction) RaiseError() uint {
	return obj.raiseError
}

// IsCondition returns true if condition, false otherwise
func (obj *instruction) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *instruction) Condition() Condition {
	return obj.condition
}

// IsAssignment returns true if assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() assignments.Assignment {
	return obj.assignment
}

// IsAccount returns true if account, false otherwise
func (obj *instruction) IsAccount() bool {
	return obj.account != nil
}

// Account returns the account, if any
func (obj *instruction) Account() accounts.Account {
	return obj.account
}

// IsDatabase returns true if database, false otherwise
func (obj *instruction) IsDatabase() bool {
	return obj.database != nil
}

// Database returns the database, if any
func (obj *instruction) Database() databases.Database {
	return obj.database
}
