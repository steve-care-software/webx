package elements

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs/kinds"
)

func TestElements_Success(t *testing.T) {
	layer := layers.NewLayerForTests(
		instructions.NewInstructionsForTests([]instructions.Instruction{
			instructions.NewInstructionWithStopForTests(),
		}),
		outputs.NewOutputForTests(
			"myVariable",
			kinds.NewKindWithContinueForTests(),
		),
		"myInput",
	)

	elements := NewElementsForTests([]Element{
		NewElementForTests(
			layer,
		),
	})

	retList := elements.List()
	if len(retList) != 1 {
		t.Errorf("the list was expected to contain 1 element")
		return
	}
}

func TestElements_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Element{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestElements_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
