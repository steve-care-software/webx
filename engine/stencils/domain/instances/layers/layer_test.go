package layers

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/outputs"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/outputs/kinds"
)

func TestLayer_Success(t *testing.T) {
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithStopForTests(),
	})

	output := outputs.NewOutputForTests(
		"myVariable",
		kinds.NewKindWithContinueForTests(),
	)

	layer := NewLayerForTests(
		instructions,
		output,
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

	if layer.HasInput() {
		t.Errorf("the layer was expected to NOT contain an input")
		return
	}
}

func TestLayer_withInput_Success(t *testing.T) {
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithStopForTests(),
	})

	output := outputs.NewOutputForTests(
		"myVariable",
		kinds.NewKindWithContinueForTests(),
	)

	input := "myInput"
	layer := NewLayerWithInputForTests(
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
