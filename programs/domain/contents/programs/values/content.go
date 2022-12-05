package values

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
	"github.com/steve-care-software/webx/domain/databases/programs/assignments"
)

type content struct {
	pInput     *uint
	assignment assignments.Assignment
	execution  entities.Identifier
	program    entities.Identifier
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
	execution entities.Identifier,
) Content {
	return createContentInternally(nil, nil, execution, nil)
}

func createContentWithProgram(
	program entities.Identifier,
) Content {
	return createContentInternally(nil, nil, nil, program)
}

func createContentInternally(
	pInput *uint,
	assignment assignments.Assignment,
	execution entities.Identifier,
	program entities.Identifier,
) Content {
	out := content{
		pInput:     pInput,
		assignment: assignment,
		execution:  execution,
		program:    program,
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
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *content) Execution() entities.Identifier {
	return obj.execution
}

// IsProgram returns true if there is a program, false otherwise
func (obj *content) IsProgram() bool {
	return obj.program != nil
}

// Program returns the program, if any
func (obj *content) Program() entities.Identifier {
	return obj.program
}
