package instructions

import "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments"

// NewInstructionsForTests creates new instructions for tests
func NewInstructionsForTests(list []Instruction) Instructions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithAssignmentForTests creates a new instruction with assignment for tests
func NewInstructionWithAssignmentForTests(assignment assignments.Assignment) Instruction {
	ins, err := NewInstructionBuilder().Create().WithAssignment(assignment).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithConditionForTests creates a new instruction with condition for tests
func NewInstructionWithConditionForTests(condition Condition) Instruction {
	ins, err := NewInstructionBuilder().Create().WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithRaiseErrorForTests creates a new instruction with raiseError for tests
func NewInstructionWithRaiseErrorForTests(raiseError uint) Instruction {
	ins, err := NewInstructionBuilder().Create().WithRaiseError(raiseError).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithStopForTests creates a new instruction with stop for tests
func NewInstructionWithStopForTests() Instruction {
	ins, err := NewInstructionBuilder().Create().IsStop().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionForTest creates a new condition for tests
func NewConditionForTest(variable string, instructions Instructions) Condition {
	ins, err := NewConditionBuilder().Create().WithVariable(variable).WithInstructions(instructions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
