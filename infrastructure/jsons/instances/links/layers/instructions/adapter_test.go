package instructions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases"
	databases_inserts "github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/inserts"
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
		instructions.NewInstructionWithAccountForTests(
			accounts.NewAccountWithInsertForTests(
				inserts.NewInsertForTests("myUser", "myPass"),
			),
		),
		instructions.NewInstructionWithDatabaseForTests(
			databases.NewDatabaseWithInsertForTests(
				databases_inserts.NewInsertForTests(
					"myContext",
					"myInstance",
					"myPath",
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
