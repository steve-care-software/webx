package instructions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/inserts"
)

func TestAdapter_Success(t *testing.T) {
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
		instructions.NewInstructionWithDatabaseForTests(
			databases.NewDatabaseWithSaveForTests(
				"mySave",
			),
		),
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
