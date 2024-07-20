package layers

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/outputs"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/outputs/kinds"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/references"
)

func TestAdapter_Success(t *testing.T) {
	ins := layers.NewLayerForTests(
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

func TestAdapter_withReferences_Success(t *testing.T) {
	ins := layers.NewLayerWithReferencesForTests(
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
		references.NewReferencesForTests([]references.Reference{
			references.NewReferenceForTests(
				"myVariable",
				[]string{"this", "is", "a", "path"},
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
