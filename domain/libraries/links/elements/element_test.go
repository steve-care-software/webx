package elements

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions/resources"
)

func TestElement_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	element := NewElementForTests(*pLayer)
	retLayer := element.Layer()
	if !bytes.Equal(pLayer.Bytes(), retLayer.Bytes()) {
		t.Errorf("the returned layer is invalid")
		return
	}

	if element.HasCondition() {
		t.Errorf("the element was expected to NOT contain condition")
		return
	}
}

func TestElement_withCondition_Success(t *testing.T) {
	condition := conditions.NewConditionForTests(
		resources.NewResourceForTests(23),
	)

	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	element := NewElementWithConditionForTests(*pLayer, condition)
	retLayer := element.Layer()
	if !bytes.Equal(pLayer.Bytes(), retLayer.Bytes()) {
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
