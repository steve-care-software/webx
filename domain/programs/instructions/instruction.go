package instructions

import "github.com/steve-care-software/syntax/domain/programs/instructions/applications"

type instruction struct {
	assignment applications.Assignment
	execution  applications.Application
	delete     applications.Application
	setPath    string
}

func createInstructionWithAssignment(
	assignment applications.Assignment,
) Instruction {
	return createInstructionInternally(assignment, nil, nil, "")
}

func createInstructionWithExecution(
	execution applications.Application,
) Instruction {
	return createInstructionInternally(nil, execution, nil, "")
}

func createInstructionWithDelete(
	delete applications.Application,
) Instruction {
	return createInstructionInternally(nil, nil, delete, "")
}

func createInstructionWithSetPath(
	setPath string,
) Instruction {
	return createInstructionInternally(nil, nil, nil, setPath)
}

func createInstructionInternally(
	assignment applications.Assignment,
	execution applications.Application,
	delete applications.Application,
	setPath string,
) Instruction {
	out := instruction{
		assignment: assignment,
		execution:  execution,
		delete:     delete,
		setPath:    setPath,
	}

	return &out
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() applications.Assignment {
	return obj.assignment
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *instruction) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *instruction) Execution() applications.Application {
	return obj.execution
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *instruction) IsDelete() bool {
	return obj.delete != nil
}

// Delete returns the delete, if any
func (obj *instruction) Delete() applications.Application {
	return obj.delete
}

// IsSetPath returns true if there is a setPath, false otherwise
func (obj *instruction) IsSetPath() bool {
	return obj.setPath != ""
}

// SetPath returns the setPath, if any
func (obj *instruction) SetPath() string {
	return obj.setPath
}
