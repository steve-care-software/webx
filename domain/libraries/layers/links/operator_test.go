package links

import "testing"

func TestOperator_isAnd_Success(t *testing.T) {
	operator := NewOperatorWithAndForTests()
	if !operator.IsAnd() {
		t.Errorf("the operator was expected to be and")
		return
	}

	if operator.IsOr() {
		t.Errorf("the operator was expected to NOT be or")
		return
	}

	if operator.IsXor() {
		t.Errorf("the operator was expected to NOT be xor")
		return
	}
}

func TestOperator_isOr_Success(t *testing.T) {
	operator := NewOperatorWithOrForTests()
	if operator.IsAnd() {
		t.Errorf("the operator was expected to NOT be and")
		return
	}

	if !operator.IsOr() {
		t.Errorf("the operator was expected to be or")
		return
	}

	if operator.IsXor() {
		t.Errorf("the operator was expected to NOT be xor")
		return
	}
}

func TestOperator_isXor_Success(t *testing.T) {
	operator := NewOperatorWithXOrForTests()
	if operator.IsAnd() {
		t.Errorf("the operator was expected to NOT be and")
		return
	}

	if operator.IsOr() {
		t.Errorf("the operator was expected to NOT be or")
		return
	}

	if !operator.IsXor() {
		t.Errorf("the operator was expected to be xor")
		return
	}
}

func TestOperator_withoutParam_returnsError(t *testing.T) {
	_, err := NewOperatorBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
