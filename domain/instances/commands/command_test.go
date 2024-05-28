package commands

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	commands_outputs "github.com/steve-care-software/datastencil/domain/instances/commands/results/outputs"
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

func TestCommand_Success(t *testing.T) {
	input := []byte("this is the command input")
	layer := layers.NewLayerForTests(
		instructions.NewInstructionsForTests([]instructions.Instruction{
			instructions.NewInstructionWithStopForTests(),
		}),
		outputs.NewOutputForTests(
			"myVariable",
			kinds.NewKindWithContinueForTests(),
		),
		"someInput",
	)

	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			commands_outputs.NewOutputForTests([]byte("this is some bytes")),
			kinds.NewKindWithPromptForTests(),
		),
	)

	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))

	parent := NewLinkWithCommandForTests(
		[]byte("this is an input"),
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
	)

	ins := NewCommandForTests(
		input,
		layer,
		result,
		parent,
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
	)

	retInput := ins.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}

	retLayer := ins.Layer()
	if !reflect.DeepEqual(layer, retLayer) {
		t.Errorf("the returned layer is invalid")
		return
	}

	retResult := ins.Result()
	if !reflect.DeepEqual(result, retResult) {
		t.Errorf("the returned result is invalid")
		return
	}

	retParent := ins.Parent()
	if !reflect.DeepEqual(parent, retParent) {
		t.Errorf("the returned parent Link is invalid")
		return
	}
}

func TestCommand_withoutParent_returnsError(t *testing.T) {
	input := []byte("this is the command input")
	layer := layers.NewLayerForTests(
		instructions.NewInstructionsForTests([]instructions.Instruction{
			instructions.NewInstructionWithStopForTests(),
		}),
		outputs.NewOutputForTests(
			"myVariable",
			kinds.NewKindWithContinueForTests(),
		),
		"someInput",
	)

	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			commands_outputs.NewOutputForTests([]byte("this is some bytes")),
			kinds.NewKindWithPromptForTests(),
		),
	)

	_, err := NewCommandBuilder().Create().WithInput(input).WithLayer(layer).WithResult(result).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommand_withoutInput_returnsError(t *testing.T) {
	layer := layers.NewLayerForTests(
		instructions.NewInstructionsForTests([]instructions.Instruction{
			instructions.NewInstructionWithStopForTests(),
		}),
		outputs.NewOutputForTests(
			"myVariable",
			kinds.NewKindWithContinueForTests(),
		),
		"someInput",
	)

	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			commands_outputs.NewOutputForTests([]byte("this is some bytes")),
			kinds.NewKindWithPromptForTests(),
		),
	)

	_, err := NewCommandBuilder().Create().WithLayer(layer).WithResult(result).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommand_withEmptyInput_returnsError(t *testing.T) {
	layer := layers.NewLayerForTests(
		instructions.NewInstructionsForTests([]instructions.Instruction{
			instructions.NewInstructionWithStopForTests(),
		}),
		outputs.NewOutputForTests(
			"myVariable",
			kinds.NewKindWithContinueForTests(),
		),
		"someInput",
	)

	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			commands_outputs.NewOutputForTests([]byte("this is some bytes")),
			kinds.NewKindWithPromptForTests(),
		),
	)

	_, err := NewCommandBuilder().Create().WithInput([]byte{}).WithLayer(layer).WithResult(result).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommand_withoutLayer_returnsError(t *testing.T) {
	input := []byte("this is the command input")
	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			commands_outputs.NewOutputForTests([]byte("this is some bytes")),
			kinds.NewKindWithPromptForTests(),
		),
	)

	_, err := NewCommandBuilder().Create().WithInput(input).WithResult(result).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommand_withoutResult_returnsError(t *testing.T) {
	input := []byte("this is the command input")
	layer := layers.NewLayerForTests(
		instructions.NewInstructionsForTests([]instructions.Instruction{
			instructions.NewInstructionWithStopForTests(),
		}),
		outputs.NewOutputForTests(
			"myVariable",
			kinds.NewKindWithContinueForTests(),
		),
		"someInput",
	)

	_, err := NewCommandBuilder().Create().WithInput(input).WithLayer(layer).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
