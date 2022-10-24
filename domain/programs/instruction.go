package programs

type instruction struct {
	assignment Assignment
	execution  Application
}

func createInstructionWithAssignment(
	assignment Assignment,
) Instruction {
	return createInstructionInternally(assignment, nil)
}

func createInstructionWithExecution(
	execution Application,
) Instruction {
	return createInstructionInternally(nil, execution)
}

func createInstructionInternally(
	assignment Assignment,
	execution Application,
) Instruction {
	out := instruction{
		assignment: assignment,
		execution:  execution,
	}

	return &out
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() Assignment {
	return obj.assignment
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *instruction) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *instruction) Execution() Application {
	return obj.execution
}
