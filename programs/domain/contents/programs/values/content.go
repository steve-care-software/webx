package values

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/contents/programs/assignments"
)

type content struct {
	pInput     *uint
	assignment assignments.Assignment
	pExecution *hash.Hash
	pProgram   *hash.Hash
}

func createContentWithInput(
	pInput *uint,
) Content {
	return createContentInternally(pInput, nil, nil, nil)
}

func createContentWithAssignment(
	assignment assignments.Assignment,
) Content {
	return createContentInternally(nil, assignment, nil, nil)
}

func createContentWithExecution(
	pExecution *hash.Hash,
) Content {
	return createContentInternally(nil, nil, pExecution, nil)
}

func createContentWithProgram(
	pProgram *hash.Hash,
) Content {
	return createContentInternally(nil, nil, nil, pProgram)
}

func createContentInternally(
	pInput *uint,
	assignment assignments.Assignment,
	pExecution *hash.Hash,
	pProgram *hash.Hash,
) Content {
	out := content{
		pInput:     pInput,
		assignment: assignment,
		pExecution: pExecution,
		pProgram:   pProgram,
	}

	return &out
}

// IsInput returns true if there is an input, false otherwise
func (obj *content) IsInput() bool {
	return obj.pInput != nil
}

// Input returns the input, if any
func (obj *content) Input() *uint {
	return obj.pInput
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *content) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *content) Assignment() assignments.Assignment {
	return obj.assignment
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *content) IsExecution() bool {
	return obj.pExecution != nil
}

// Execution returns the execution, if any
func (obj *content) Execution() *hash.Hash {
	return obj.pExecution
}

// IsProgram returns true if there is a program, false otherwise
func (obj *content) IsProgram() bool {
	return obj.pProgram != nil
}

// Program returns the program, if any
func (obj *content) Program() *hash.Hash {
	return obj.pProgram
}
