package instructions

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/contents/programs/assignments"
)

type content struct {
	assignment assignments.Assignment
	execution  hash.Hash
}

func createContentWithAssignment(
	assignment assignments.Assignment,
) Content {
	return createContentInternally(assignment, nil)
}

func createContentWithExecution(
	execution hash.Hash,
) Content {
	return createContentInternally(nil, execution)
}

func createContentInternally(
	assignment assignments.Assignment,
	execution hash.Hash,
) Content {
	out := content{
		assignment: assignment,
		execution:  execution,
	}

	return &out
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
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *content) Execution() hash.Hash {
	return obj.execution
}
