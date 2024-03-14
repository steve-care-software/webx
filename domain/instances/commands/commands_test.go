package commands

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins/resources"
)

func TestCommands_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))

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
					results.NewOutputForTests([]byte("this is some bytes")),
					kinds.NewKindWithPromptForTests(),
				),
			),
			NewLinkForTests(
				[]byte("some input"),
				links.NewLinkForTests(
					origins.NewOriginForTests(
						resources.NewResourceForTests(*pFirstLayer),
						operators.NewOperatorWithAndForTests(),
						origins.NewValueWithResourceForTests(
							resources.NewResourceForTests(*pSecondLayer),
						),
					),
					elements.NewElementsForTests([]elements.Element{
						elements.NewElementForTests(*pLayer),
					}),
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
