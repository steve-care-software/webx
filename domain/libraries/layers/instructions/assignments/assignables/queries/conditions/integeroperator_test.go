package conditions

import (
	"testing"
)

func TestIntegerOperator_isSmallerThan_isEqual_Success(t *testing.T) {
	ins := NewIntegerOperatorWithIsSmallerThanAndIsEqualForTests()

	if !ins.IsSmallerThan() {
		t.Errorf("the integerOperator was expected to contain a isSmallerThan")
		return
	}

	if ins.IsBiggerThan() {
		t.Errorf("the integerOperator was expected to NOT contain a isBiggerThan")
		return
	}

	if !ins.IsEqual() {
		t.Errorf("the integerOperator was expected to contain a isEqual")
		return
	}
}

func TestIntegerOperator_isSmallerThan_Success(t *testing.T) {
	ins := NewIntegerOperatorWithIsSmallerThanForTests()

	if !ins.IsSmallerThan() {
		t.Errorf("the integerOperator was expected to contain a isSmallerThan")
		return
	}

	if ins.IsBiggerThan() {
		t.Errorf("the integerOperator was expected to NOT contain a isBiggerThan")
		return
	}

	if ins.IsEqual() {
		t.Errorf("the integerOperator was expected to NOT contain a isEqual")
		return
	}
}

func TestIntegerOperator_isBiggerThan_isEqual_Success(t *testing.T) {
	ins := NewIntegerOperatorWithIsBiggerThanAndIsEqualForTests()

	if ins.IsSmallerThan() {
		t.Errorf("the integerOperator was expected to NOT contain a isSmallerThan")
		return
	}

	if !ins.IsBiggerThan() {
		t.Errorf("the integerOperator was expected to contain a isBiggerThan")
		return
	}

	if !ins.IsEqual() {
		t.Errorf("the integerOperator was expected to contain a isEqual")
		return
	}
}

func TestIntegerOperator_isBiggerThan_Success(t *testing.T) {
	ins := NewIntegerOperatorWithIsBiggerThanForTests()

	if ins.IsSmallerThan() {
		t.Errorf("the integerOperator was expected to NOT contain a isSmallerThan")
		return
	}

	if !ins.IsBiggerThan() {
		t.Errorf("the integerOperator was expected to contain a isBiggerThan")
		return
	}

	if ins.IsEqual() {
		t.Errorf("the integerOperator was expected to NOT contain a isEqual")
		return
	}
}

func TestIntegerOperator_isEqual_Success(t *testing.T) {
	ins := NewIntegerOperatorWithIsEqualForTests()

	if ins.IsSmallerThan() {
		t.Errorf("the integerOperator was expected to NOT contain a isSmallerThan")
		return
	}

	if ins.IsBiggerThan() {
		t.Errorf("the integerOperator was expected to NOT contain a isBiggerThan")
		return
	}

	if !ins.IsEqual() {
		t.Errorf("the integerOperator was expected to contain a isEqual")
		return
	}
}

func TestIntegerOperator_withoutParam_returnsError(t *testing.T) {
	_, err := NewIntegerOperatorBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
