package elements

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions/resources"
)

func TestElement_Success(t *testing.T) {
	layer := []string{"this", "is", "a", "path"}
	element := NewElementForTests(
		layer,
	)

	retLayer := element.Layer()
	if !reflect.DeepEqual(layer, retLayer) {
		t.Errorf("the returned layer is invalid")
		return
	}

	if element.HasCondition() {
		t.Errorf("the element was expected to NOT contain condition")
		return
	}
}

func TestElement_withCondition_Success(t *testing.T) {
	layer := []string{"this", "is", "a", "path"}
	condition := conditions.NewConditionForTests(
		resources.NewResourceForTests(23),
	)

	element := NewElementWithConditionForTests(layer, condition)
	retLayer := element.Layer()
	if !reflect.DeepEqual(layer, retLayer) {
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
