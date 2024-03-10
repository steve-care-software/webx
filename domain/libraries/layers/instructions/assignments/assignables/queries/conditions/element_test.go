package conditions

import (
	"reflect"
	"testing"
)

func TestElement_withResource_Success(t *testing.T) {
	resource := NewResourceWithFieldForTests(
		NewPointerForTests("myEntity", "myField"),
	)

	ins := NewElementWithResourceForTests(resource)

	if !ins.IsResource() {
		t.Errorf("the element was expected to contain a resource")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the element was expected to NOT contain a condition")
		return
	}

	retResource := ins.Resource()
	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}
}

func TestElement_withCondition_Success(t *testing.T) {
	condition := NewConditionForTests(
		NewPointerForTests("myEntity", "myField"),
		NewOperatorWithEqualForTests(),
		NewElementWithResourceForTests(
			NewResourceWithValueForTests(45),
		),
	)

	ins := NewElementWithConditionForTests(condition)

	if ins.IsResource() {
		t.Errorf("the element was expected to NOT contain a resource")
		return
	}

	if !ins.IsCondition() {
		t.Errorf("the element was expected to contain a condition")
		return
	}

	retCondition := ins.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the condition is invalid")
		return
	}
}

func TestElement_withoutParam_returnsError(t *testing.T) {
	_, err := NewElementBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
