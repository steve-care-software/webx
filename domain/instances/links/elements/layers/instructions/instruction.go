package instructions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists"
)

type instruction struct {
	hash       hash.Hash
	isStop     bool
	raiseError uint
	condition  Condition
	assignment assignments.Assignment
	database   databases.Database
	list       lists.List
	loop       Loop
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
		database,
		nil,
		nil,
	)
}

func createInstructionWithList(
	hash hash.Hash,
	list lists.List,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		nil,
		nil,
		list,
		nil,
	)
}

func createInstructionWithLoop(
	hash hash.Hash,
	loop Loop,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		nil,
		nil,
		nil,
		loop,
	)
}

func createInstructionInternally(
	hash hash.Hash,
	isStop bool,
	raiseError uint,
	condition Condition,
	assignment assignments.Assignment,
	database databases.Database,
	list lists.List,
	loop Loop,
) Instruction {
	out := instruction{
		hash:       hash,
		isStop:     isStop,
		raiseError: raiseError,
		condition:  condition,
		assignment: assignment,
		database:   database,
		list:       list,
		loop:       loop,
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

// IsDatabase returns true if database, false otherwise
func (obj *instruction) IsDatabase() bool {
	return obj.database != nil
}

// Database returns the database, if any
func (obj *instruction) Database() databases.Database {
	return obj.database
}

// IsList returns true if list, false otherwise
func (obj *instruction) IsList() bool {
	return obj.list != nil
}

// List returns the list, if any
func (obj *instruction) List() lists.List {
	return obj.list
}

// IsLoop returns true if loop, false otherwise
func (obj *instruction) IsLoop() bool {
	return obj.loop != nil
}

// Loop returns the loop, if any
func (obj *instruction) Loop() Loop {
	return obj.loop
}
