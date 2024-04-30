package relationals

import (
	"testing"
)

func TestRelational_isAnd_Success(t *testing.T) {
	ins := NewRelationalWithAndForTests()

	if !ins.IsAnd() {
		t.Errorf("the relational was expected to contain a isAnd")
		return
	}

	if ins.IsOr() {
		t.Errorf("the relational was expected to NOT contain a isOr")
		return
	}
}

func TestRelational_isor_Success(t *testing.T) {
	ins := NewRelationalWithOrForTests()

	if ins.IsAnd() {
		t.Errorf("the relational was expected to NOT contain a isAnd")
		return
	}

	if !ins.IsOr() {
		t.Errorf("the relational was expected to contain a isOr")
		return
	}
}

func TestRelational_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
