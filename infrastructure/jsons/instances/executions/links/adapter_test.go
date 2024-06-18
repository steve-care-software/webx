package links

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/success"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/success/outputs"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/bytes"
	source_layers_outputs "github.com/steve-care-software/datastencil/domain/instances/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs/kinds"
	source_links "github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions/resources"
)

func TestAdapter_Success(t *testing.T) {
	ins := links.NewLinkForTests(
		[]byte("this is some link input"),
		source_links.NewLinkForTests(
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementWithConditionForTests(
					[]string{"path", "to", "layer"},
					conditions.NewConditionForTests(
						resources.NewResourceForTests(uint(45)),
					),
				),
				elements.NewElementForTests(
					[]string{"another", "path", "to", "layer"},
				),
			}),
		),
	)

	adapter := NewAdapter()

	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withLayers_Success(t *testing.T) {
	ins := links.NewLinkWithLayersForTests(
		[]byte("this is some link input"),
		source_links.NewLinkForTests(
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementWithConditionForTests(
					[]string{"path", "to", "layer"},
					conditions.NewConditionForTests(
						resources.NewResourceForTests(uint(45)),
					),
				),
				elements.NewElementForTests(
					[]string{"another", "path", "to", "layer"},
				),
			}),
		),
		layers.NewLayersForTests([]layers.Layer{
			layers.NewLayerForTests(
				[]byte("this is some input"),
				source_layers.NewLayerForTests(
					instructions.NewInstructionsForTests([]instructions.Instruction{
						instructions.NewInstructionWithAssignmentForTests(
							assignments.NewAssignmentForTests(
								"anotherName",
								assignables.NewAssignableWithBytesForTests(
									bytes_domain.NewBytesWithHashBytesForTests(
										"anotherInput",
									),
								),
							),
						),
						instructions.NewInstructionWithRaiseErrorForTests(22),
						instructions.NewInstructionWithStopForTests(),
					}),
					source_layers_outputs.NewOutputForTests(
						"myVariable",
						kinds.NewKindWithContinueForTests(),
					),
					"myInput",
				),
				results.NewResultWithSuccessForTests(
					success.NewSuccessForTests(
						outputs.NewOutputForTests(
							[]byte("this is an input"),
						),
						kinds.NewKindWithPromptForTests(),
					),
				),
			),
		}),
	)

	adapter := NewAdapter()

	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
