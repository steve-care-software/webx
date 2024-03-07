package layers

import (
	"reflect"
	"testing"
)

func TestCondition_Success(t *testing.T) {
	variable := "myName"
	instructions := NewInstructionsForTests([]Instruction{
		NewInstructionWithStopForTests(),
	})

	condition := NewConditionForTest(variable, instructions)
	retVariable := condition.Variable()
	if variable != retVariable {
		t.Errorf("the variable was expected to be '%s', '%s' returned", variable, retVariable)
		return
	}

	retInstructions := condition.Instructions()
	if !reflect.DeepEqual(instructions, retInstructions) {
		t.Errorf("the instructions is invalid")
		return
	}
}

func TestCondition_withoutVariable_returnsError(t *testing.T) {
	instructions := NewInstructionsForTests([]Instruction{
		NewInstructionWithStopForTests(),
	})

	_, err := NewConditionBuilder().Create().WithInstructions(instructions).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCondition_withoutInstructions_returnsError(t *testing.T) {
	variable := "myName"
	_, err := NewConditionBuilder().Create().WithVariable(variable).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
