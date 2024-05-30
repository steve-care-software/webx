package layers

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs/kinds"
)

func TestLayer_Success(t *testing.T) {
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithStopForTests(),
	})

	output := outputs.NewOutputForTests(
		"myVariable",
		kinds.NewKindWithContinueForTests(),
	)

	input := "myInput"
	layer := NewLayerForTests(
		instructions,
		output,
		input,
	)

	retInstructions := layer.Instructions()
	if !reflect.DeepEqual(instructions, retInstructions) {
		t.Errorf("the returned instructions is invalid")
		return
	}

	retOutput := layer.Output()
	if !reflect.DeepEqual(output, retOutput) {
		t.Errorf("the returned output is invalid")
		return
	}

	retInput := layer.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}
}

func TestLayer_withoutInstructions_returnsError(t *testing.T) {
	output := outputs.NewOutputForTests(
		"myVariable",
		kinds.NewKindWithContinueForTests(),
	)

	input := "myInput"
	_, err := NewBuilder().Create().WithOutput(output).WithInput(input).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestLayer_withoutOutput_returnsError(t *testing.T) {
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithStopForTests(),
	})

	input := "myInput"
	_, err := NewBuilder().Create().WithInstructions(instructions).WithInput(input).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestLayer_withoutInput_returnsError(t *testing.T) {
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithStopForTests(),
	})

	output := outputs.NewOutputForTests(
		"myVariable",
		kinds.NewKindWithContinueForTests(),
	)
	_, err := NewBuilder().Create().WithInstructions(instructions).WithOutput(output).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
