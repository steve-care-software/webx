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

func TestLink_withCommand_Success(t *testing.T) {
	input := []byte("this is an input")

	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	link := links.NewLinkForTests(
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
	)

	command := NewCommandForTests(
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
	)

	ins := NewLinkWithCommandForTests(input, link, command)

	retInput := ins.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}

	retLink := ins.Link()
	if !reflect.DeepEqual(link, retLink) {
		t.Errorf("the returned link is invalid")
		return
	}

	retCommand := ins.Command()
	if !reflect.DeepEqual(command, retCommand) {
		t.Errorf("the returned command is invalid")
		return
	}
}

func TestLink_withoutInput_returnsError(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	link := links.NewLinkForTests(
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
	)

	command := NewCommandForTests(
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
	)

	_, err := NewLinkBuilder().Create().WithLink(link).WithCommand(command).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestLink_withoutLink_returnsError(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))

	input := []byte("this is an input")
	command := NewCommandForTests(
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
	)

	_, err := NewLinkBuilder().Create().WithInput(input).WithCommand(command).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
