package links

import (
	"reflect"
	"testing"
)

func TestCondition_Success(t *testing.T) {
	resource := NewConditionResourceForTests(23)
	condition := NewConditionForTests(
		resource,
	)

	if condition.HasNext() {
		t.Errorf("the condition was expected to NOT contain a next")
		return
	}

	retResource := condition.Resource()
	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}
}

func TestCondition_withNext_Success(t *testing.T) {
	resource := NewConditionResourceForTests(23)
	next := NewConditionValueWithResourceForTests(
		NewConditionResourceForTests(44),
	)

	condition := NewConditionWithNextForTests(
		resource,
		next,
	)

	if !condition.HasNext() {
		t.Errorf("the condition was expected to contain a next")
		return
	}

	retResource := condition.Resource()
	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}

	retNext := condition.Next()
	if !reflect.DeepEqual(next, retNext) {
		t.Errorf("the next is invalid")
		return
	}
}
func TestCondition__withoutResource_returnsError(t *testing.T) {
	_, err := NewConditionBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
