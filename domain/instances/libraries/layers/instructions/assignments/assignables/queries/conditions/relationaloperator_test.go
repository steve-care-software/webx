package conditions

import (
	"testing"
)

func TestRelationalOperator_isAnd_Success(t *testing.T) {
	ins := NewRelationalOperatorWithAndForTests()

	if !ins.IsAnd() {
		t.Errorf("the relationalOperator was expected to contain a isAnd")
		return
	}

	if ins.IsOr() {
		t.Errorf("the relationalOperator was expected to NOT contain a isOr")
		return
	}
}

func TestRelationalOperator_isor_Success(t *testing.T) {
	ins := NewRelationalOperatorWithOrForTests()

	if ins.IsAnd() {
		t.Errorf("the relationalOperator was expected to NOT contain a isAnd")
		return
	}

	if !ins.IsOr() {
		t.Errorf("the relationalOperator was expected to contain a isOr")
		return
	}
}

func TestRelationalOperator_withoutParam_returnsError(t *testing.T) {
	_, err := NewRelationalOperatorBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
