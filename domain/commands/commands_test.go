package commands

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/commands/results"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/outputs/kinds"
)

func TestCommands_Success(t *testing.T) {
	list := []Command{
		NewCommandForTests(
			[]byte("this is the command input"),
			layers.NewLayerForTests(
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithStopForTests(),
				}),
				outputs.NewOutputForTests(
					"myVariable",
					kinds.NewKindWithContinueForTests(),
				),
				"someInput",
			),
			results.NewResultWithSuccessForTests(
				results.NewSuccessForTests(
					[]byte("this is some bytes"),
					kinds.NewKindWithPromptForTests(),
				),
			),
		),
	}

	ins := NewCommandsForTests(list)

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestCommands_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommands_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Command{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
