package commands

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

func TestLink_withCommand_Success(t *testing.T) {
	input := []byte("this is an input")
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))

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
			elements.NewElementForTests(*pLayer),
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
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))

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
			elements.NewElementForTests(*pLayer),
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
	)

	_, err := NewLinkBuilder().Create().WithLink(link).WithCommand(command).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestLink_withoutLink_returnsError(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
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
	)

	_, err := NewLinkBuilder().Create().WithInput(input).WithCommand(command).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
