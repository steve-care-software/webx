package layers

import (
	"reflect"
	"testing"
)

func TestInstruction_isStop_Success(t *testing.T) {
	instruction := NewInstructionWithStopForTests()
	if !instruction.IsStop() {
		t.Errorf("the instruction was expected to contain a stop")
		return
	}

	if instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to NOT contain a raiseError")
		return
	}

	if instruction.IsCondition() {
		t.Errorf("the instruction was expected to NOT contain a condition")
		return
	}

	if instruction.IsAssignment() {
		t.Errorf("the instruction was expected to NOT contain an assignment")
		return
	}
}

func TestInstruction_withRaiseError_Success(t *testing.T) {
	code := uint(56)
	instruction := NewInstructionWithRaiseErrorForTests(code)
	if instruction.IsStop() {
		t.Errorf("the instruction was expected to NOT contain a stop")
		return
	}

	if !instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to contain a raiseError")
		return
	}

	if instruction.IsCondition() {
		t.Errorf("the instruction was expected to NOT contain a condition")
		return
	}

	if instruction.IsAssignment() {
		t.Errorf("the instruction was expected to NOT contain an assignment")
		return
	}

	retRaiseError := instruction.RaiseError()
	if code != retRaiseError {
		t.Errorf("the raisedError code was expected to be %d, %d returned", code, retRaiseError)
		return
	}
}

func TestInstruction_withCondition_Success(t *testing.T) {
	condition := NewConditionForTest(
		"myName",
		NewInstructionsForTests([]Instruction{
			NewInstructionWithStopForTests(),
		}),
	)

	instruction := NewInstructionWithConditionForTests(condition)
	if instruction.IsStop() {
		t.Errorf("the instruction was expected to NOT contain a stop")
		return
	}

	if instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to NOT contain a raiseError")
		return
	}

	if !instruction.IsCondition() {
		t.Errorf("the instruction was expected to contain a condition")
		return
	}

	if instruction.IsAssignment() {
		t.Errorf("the instruction was expected to NOT contain an assignment")
		return
	}

	retCondition := instruction.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the returned condition is invalid")
		return
	}
}

func TestInstruction_withAssignment_Success(t *testing.T) {
	assignment := NewAssignmentForTests(
		"myName",
		NewAssignableWithBytesForTests(NewBytesWithJoinForTests([]string{
			"first",
			"second",
		})),
	)

	instruction := NewInstructionWithAssignmentForTests(assignment)
	if instruction.IsStop() {
		t.Errorf("the instruction was expected to NOT contain a stop")
		return
	}

	if instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to NOT contain a raiseError")
		return
	}

	if instruction.IsCondition() {
		t.Errorf("the instruction was expected to NOT contain a condition")
		return
	}

	if !instruction.IsAssignment() {
		t.Errorf("the instruction was expected to contain an assignment")
		return
	}

	retAssignment := instruction.Assignment()
	if !reflect.DeepEqual(assignment, retAssignment) {
		t.Errorf("the returned assignment is invalid")
		return
	}
}

func TestInstruction_withoutParam_returnsError(t *testing.T) {
	_, err := NewInstructionBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
