package elements

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions/resources"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs/kinds"
)

func TestElement_Success(t *testing.T) {
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

	element := NewElementForTests(
		layer,
	)

	retLayer := element.Layer()
	if !bytes.Equal(layer.Hash().Bytes(), retLayer.Hash().Bytes()) {
		t.Errorf("the returned layer is invalid")
		return
	}

	if element.HasCondition() {
		t.Errorf("the element was expected to NOT contain condition")
		return
	}
}

func TestElement_withCondition_Success(t *testing.T) {
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

	condition := conditions.NewConditionForTests(
		resources.NewResourceForTests(23),
	)

	element := NewElementWithConditionForTests(layer, condition)
	retLayer := element.Layer()
	if !bytes.Equal(layer.Hash().Bytes(), retLayer.Hash().Bytes()) {
		t.Errorf("the returned layer is invalid")
		return
	}

	if !element.HasCondition() {
		t.Errorf("the element was expected to contain condition")
		return
	}

	retCondition := element.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the condition is invalid")
		return
	}
}
