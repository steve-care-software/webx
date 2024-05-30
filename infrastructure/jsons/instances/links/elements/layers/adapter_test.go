package layers

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs/kinds"
)

func TestAdapter_Success(t *testing.T) {
	ins := layers.NewLayersForTests([]layers.Layer{
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
