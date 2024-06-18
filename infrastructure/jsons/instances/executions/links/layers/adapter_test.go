package layers

import (
	"bytes"
	"testing"

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
)

func TestAdapter_Success(t *testing.T) {

	ins := layers.NewLayersForTests([]layers.Layer{
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
	})

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
