package conditions

import (
	"reflect"
	"testing"
)

func TestOperator_isEqual_Success(t *testing.T) {
	ins := NewOperatorWithEqualForTests()

	if !ins.IsEqual() {
		t.Errorf("the operator was expected to contain a isEqual")
		return
	}

	if ins.IsRelational() {
		t.Errorf("the operator was expected to NOT contain a relational")
		return
	}

	if ins.IsInteger() {
		t.Errorf("the operator was expected to NOT contain an integer")
		return
	}
}

func TestOperator_isRelational_Success(t *testing.T) {
	relational := NewRelationalOperatorWithAndForTests()
	ins := NewOperatorWithRelationalForTests(
		relational,
	)

	if ins.IsEqual() {
		t.Errorf("the operator was expected to NOT contain a isEqual")
		return
	}

	if !ins.IsRelational() {
		t.Errorf("the operator was expected to contain a relational")
		return
	}

	if ins.IsInteger() {
		t.Errorf("the operator was expected to NOT contain an integer")
		return
	}

	retRelational := ins.Relational()
	if !reflect.DeepEqual(relational, retRelational) {
		t.Errorf("the relational is invalid")
		return
	}
}

func TestOperator_isInteger_Success(t *testing.T) {
	integer := NewIntegerOperatorWithIsSmallerThanForTests()
	ins := NewOperatorWithIntegerForTests(
		integer,
	)

	if ins.IsEqual() {
		t.Errorf("the operator was expected to NOT contain a isEqual")
		return
	}

	if ins.IsRelational() {
		t.Errorf("the operator was expected to NOT contain a relational")
		return
	}

	if !ins.IsInteger() {
		t.Errorf("the operator was expected to contain an integer")
		return
	}

	retInteger := ins.Integer()
	if !reflect.DeepEqual(integer, retInteger) {
		t.Errorf("the integer is invalid")
		return
	}
}

func TestOperator_withoutParam_returnsError(t *testing.T) {
	_, err := NewOperatorBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
