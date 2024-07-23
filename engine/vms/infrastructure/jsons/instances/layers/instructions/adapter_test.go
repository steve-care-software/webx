package instructions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists/inserts"
)

func TestAdapter_list_Success(t *testing.T) {
	ins := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				"myName",
				assignables.NewAssignableWithBytesForTests(
					bytes_domain.NewBytesWithHashBytesForTests(
						"myInput",
					),
				),
			),
		),
		instructions.NewInstructionWithConditionForTests(
			instructions.NewConditionForTest(
				"myCondition",
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
			),
		),
		instructions.NewInstructionWithRaiseErrorForTests(33),
		instructions.NewInstructionWithStopForTests(),
		instructions.NewInstructionWithListForTests(
			lists.NewListWithInsertForTests(
				inserts.NewInsertForTests(
					"myList",
					"myElement",
				),
			),
		),
		instructions.NewInstructionWithLoopForTests(
			instructions.NewLoopForTest(
				"myAmount",
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithListForTests(
						lists.NewListWithInsertForTests(
							inserts.NewInsertForTests(
								"myList",
								"myElement",
							),
						),
					),
				}),
			),
		),
		instructions.NewInstructionWithExecutionForTests(
			executions.NewExecutionForTests(
				"myExecutable",
				executions.NewContentWithCommitForTests("myCommit"),
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

func TestAdapter_Success(t *testing.T) {
	ins := instructions.NewInstructionWithExecutionForTests(
		executions.NewExecutionForTests(
			"myExecutable",
			executions.NewContentWithCommitForTests("myCommit"),
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
