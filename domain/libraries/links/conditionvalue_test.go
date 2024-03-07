package links

import (
	"reflect"
	"testing"
)

func TestConditionValue_withResource_Success(t *testing.T) {
	conditionResource := NewConditionResourceForTests(45)
	conditionValue := NewConditionValueWithResourceForTests(conditionResource)

	if !conditionValue.IsResource() {
		t.Errorf("the conditionValue was expected to contain a resource")
		return
	}

	if conditionValue.IsCondition() {
		t.Errorf("the conditionValue was expected to NOT contain a resource")
		return
	}

	retResource := conditionValue.Resource()
	if !reflect.DeepEqual(conditionResource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}

}

func TestConditionValue_withCondition_Success(t *testing.T) {
	condition := NewConditionForTests(
		NewConditionResourceForTests(22),
	)
	conditionValue := NewConditionValueWithConditionForTests(condition)

	if conditionValue.IsResource() {
		t.Errorf("the conditionValue was expected to NOT contain a resource")
		return
	}

	if !conditionValue.IsCondition() {
		t.Errorf("the conditionValue was expected to contain a resource")
		return
	}

	retCondition := conditionValue.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the condition is invalid")
		return
	}

}

func TestConditionValue_withoutParam_returnsError(t *testing.T) {
	_, err := NewConditionValueBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
