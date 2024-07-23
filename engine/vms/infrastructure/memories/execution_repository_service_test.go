package memories

import (
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/success"
	success_outputs "github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs/kinds"
)

func TestExecution_Success(t *testing.T) {
	dbPath := []string{"my", "db", "path"}
	first := executions.NewExecutionWithInputForTests(
		[]byte("myInput"),
		layers.NewLayerWithInputForTests(
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

	repository, service := NewExecutionRepositoryAndService()
	err := service.Save(dbPath, first)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retExecutions, err := repository.RetrieveAll(dbPath, []hash.Hash{
		first.Hash(),
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retList := retExecutions.List()
	if len(retList) != 1 {
		t.Errorf("%d executions were expected, %d returned", 1, len(retList))
		return
	}
}
