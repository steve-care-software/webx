package conditions

import (
	"reflect"
	"testing"
)

func TestCondition_Success(t *testing.T) {
	pointer := NewPointerForTests("myEntity", "myField")
	operator := NewOperatorWithEqualForTests()
	element := NewElementWithResourceForTests(
		NewResourceWithValueForTests(45),
	)

	ins := NewConditionForTests(
		pointer,
		operator,
		element,
	)

	retPointer := ins.Pointer()
	if !reflect.DeepEqual(pointer, retPointer) {
		t.Errorf("the pointer is invalid")
		return
	}

	retOperator := ins.Operator()
	if !reflect.DeepEqual(operator, retOperator) {
		t.Errorf("the operator is invalid")
		return
	}

	retElement := ins.Element()
	if !reflect.DeepEqual(element, retElement) {
		t.Errorf("the element is invalid")
		return
	}
}

func TestCondition_withoutPointer_returnsError(t *testing.T) {
	operator := NewOperatorWithEqualForTests()
	element := NewElementWithResourceForTests(
		NewResourceWithValueForTests(45),
	)

	_, err := NewBuilder().Create().WithOperator(operator).WithElement(element).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestCondition_withoutOperator_returnsError(t *testing.T) {
	pointer := NewPointerForTests("myEntity", "myField")
	element := NewElementWithResourceForTests(
		NewResourceWithValueForTests(45),
	)

	_, err := NewBuilder().Create().WithPointer(pointer).WithElement(element).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestCondition_withoutElement_returnsError(t *testing.T) {
	pointer := NewPointerForTests("myEntity", "myField")
	operator := NewOperatorWithEqualForTests()

	_, err := NewBuilder().Create().WithOperator(operator).WithPointer(pointer).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
