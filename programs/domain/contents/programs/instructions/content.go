package instructions

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type content struct {
	pAssignment *hash.Hash
	pExecution  *hash.Hash
}

func createContentWithAssignment(
	pAssignment *hash.Hash,
) Content {
	return createContentInternally(pAssignment, nil)
}

func createContentWithExecution(
	pExecution *hash.Hash,
) Content {
	return createContentInternally(nil, pExecution)
}

func createContentInternally(
	pAssignment *hash.Hash,
	pExecution *hash.Hash,
) Content {
	out := content{
		pAssignment: pAssignment,
		pExecution:  pExecution,
	}

	return &out
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *content) IsAssignment() bool {
	return obj.pAssignment != nil
}

// Assignment returns the assignment, if any
func (obj *content) Assignment() *hash.Hash {
	return obj.pAssignment
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *content) IsExecution() bool {
	return obj.pExecution != nil
}

// Execution returns the execution, if any
func (obj *content) Execution() *hash.Hash {
	return obj.pExecution
}
