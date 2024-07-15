package executions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results/success"
	success_outputs "github.com/steve-care-software/datastencil/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs/kinds"
)

func TestAdapter_withExecutions_Success(t *testing.T) {

	ins := executions.NewExecutionsForTests([]executions.Execution{
		executions.NewExecutionForTests(
			[]byte("myInput"),
			layers.NewLayerForTests(
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
				outputs.NewOutputForTests(
					"myVariable",
					kinds.NewKindWithContinueForTests(),
				),
				"myInput",
			),
			results.NewResultWithSuccessForTests(
				success.NewSuccessForTests(
					success_outputs.NewOutputForTests(
						[]byte("this is an input"),
					),
					kinds.NewKindWithPromptForTests(),
				),
			),
		),
	})

	adapter := NewAdapter()

	retBytes, err := adapter.InstancesToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withExecution_Success(t *testing.T) {

	ins := executions.NewExecutionForTests(
		[]byte("myInput"),
		layers.NewLayerForTests(
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
			outputs.NewOutputForTests(
				"myVariable",
				kinds.NewKindWithContinueForTests(),
			),
			"myInput",
		),
		results.NewResultWithSuccessForTests(
			success.NewSuccessForTests(
				success_outputs.NewOutputForTests(
					[]byte("this is an input"),
				),
				kinds.NewKindWithPromptForTests(),
			),
		),
	)

	adapter := NewAdapter()

	retBytes, err := adapter.InstanceToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
