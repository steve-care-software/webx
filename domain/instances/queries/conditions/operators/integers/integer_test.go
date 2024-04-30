package integers

import (
	"testing"
)

func TestInteger_isSmallerThan_isEqual_Success(t *testing.T) {
	ins := NewIntegerWithIsSmallerThanAndIsEqualForTests()

	if !ins.IsSmallerThan() {
		t.Errorf("the integer was expected to contain a isSmallerThan")
		return
	}

	if ins.IsBiggerThan() {
		t.Errorf("the integer was expected to NOT contain a isBiggerThan")
		return
	}

	if !ins.IsEqual() {
		t.Errorf("the integer was expected to contain a isEqual")
		return
	}
}

func TestInteger_isSmallerThan_Success(t *testing.T) {
	ins := NewIntegerWithIsSmallerThanForTests()

	if !ins.IsSmallerThan() {
		t.Errorf("the integer was expected to contain a isSmallerThan")
		return
	}

	if ins.IsBiggerThan() {
		t.Errorf("the integer was expected to NOT contain a isBiggerThan")
		return
	}

	if ins.IsEqual() {
		t.Errorf("the integer was expected to NOT contain a isEqual")
		return
	}
}

func TestInteger_isBiggerThan_isEqual_Success(t *testing.T) {
	ins := NewIntegerWithIsBiggerThanAndIsEqualForTests()

	if ins.IsSmallerThan() {
		t.Errorf("the integer was expected to NOT contain a isSmallerThan")
		return
	}

	if !ins.IsBiggerThan() {
		t.Errorf("the integer was expected to contain a isBiggerThan")
		return
	}

	if !ins.IsEqual() {
		t.Errorf("the integer was expected to contain a isEqual")
		return
	}
}

func TestInteger_isBiggerThan_Success(t *testing.T) {
	ins := NewIntegerWithIsBiggerThanForTests()

	if ins.IsSmallerThan() {
		t.Errorf("the integer was expected to NOT contain a isSmallerThan")
		return
	}

	if !ins.IsBiggerThan() {
		t.Errorf("the integer was expected to contain a isBiggerThan")
		return
	}

	if ins.IsEqual() {
		t.Errorf("the integer was expected to NOT contain a isEqual")
		return
	}
}

func TestInteger_isEqual_Success(t *testing.T) {
	ins := NewIntegerWithIsEqualForTests()

	if ins.IsSmallerThan() {
		t.Errorf("the integer was expected to NOT contain a isSmallerThan")
		return
	}

	if ins.IsBiggerThan() {
		t.Errorf("the integer was expected to NOT contain a isBiggerThan")
		return
	}

	if !ins.IsEqual() {
		t.Errorf("the integer was expected to contain a isEqual")
		return
	}
}

func TestInteger_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
