package instructions

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/executions"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/lists"
)

// NewInstructionsForTests creates new instructions for tests
func NewInstructionsForTests(list []Instruction) Instructions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithExecutionForTests creates a new instruction with execution for tests
func NewInstructionWithExecutionForTests(execution executions.Execution) Instruction {
	ins, err := NewInstructionBuilder().Create().WithExecution(execution).Now()
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

// NewInstructionWithListForTests creates a new instruction with list for tests
func NewInstructionWithListForTests(list lists.List) Instruction {
	ins, err := NewInstructionBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInstructionWithLoopForTests creates a new instruction with loop for tests
func NewInstructionWithLoopForTests(loop Loop) Instruction {
	ins, err := NewInstructionBuilder().Create().WithLoop(loop).Now()
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

// NewLoopForTest creates a new loop for tests
func NewLoopForTest(amount string, instructions Instructions) Loop {
	ins, err := NewLoopBuuilder().Create().WithAmount(amount).WithInstructions(instructions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
