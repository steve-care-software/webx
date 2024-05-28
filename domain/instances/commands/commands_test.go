package commands

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/success"
	commands_outputs "github.com/steve-care-software/datastencil/domain/instances/commands/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

func TestCommands_Success(t *testing.T) {
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
				success.NewSuccessForTests(
					commands_outputs.NewOutputForTests([]byte("this is some bytes")),
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
						elements.NewElementForTests(
							layers.NewLayerForTests(
								instructions.NewInstructionsForTests([]instructions.Instruction{
									instructions.NewInstructionWithStopForTests(),
								}),
								outputs.NewOutputForTests(
									"myVariable",
									kinds.NewKindWithContinueForTests(),
								),
								"myInput",
							),
						),
					}),
				),
			),
			commits.NewCommitForTests(
				"This is a description",
				actions.NewActionsForTests([]actions.Action{
					actions.NewActionWithModificationsForTests(
						[]string{"this", "is", "a", "path"},
						modifications.NewModificationsForTests([]modifications.Modification{
							modifications.NewModificationWithInsertForTests([]byte("some data to insert")),
							modifications.NewModificationWithDeleteForTests(
								deletes.NewDeleteForTests(
									0,
									50,
								),
							),
						}),
					),
				}),
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
