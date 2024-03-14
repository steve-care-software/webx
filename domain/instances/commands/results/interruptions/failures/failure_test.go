package failures

import (
	"testing"
)

func TestFailure_Success(t *testing.T) {
	index := uint(32)
	code := uint(56)
	isRaisedInLayer := false
	ins := NewFailureForTests(index, code, isRaisedInLayer)
	retCode := ins.Code()
	if code != retCode {
		t.Errorf("the code was expected to be %d, %d returned", code, retCode)
		return
	}

	retIndex := ins.Index()
	if index != retIndex {
		t.Errorf("the index was expected to be %d, %d returned", index, retIndex)
		return
	}

	retIsRaisedInLayer := ins.IsRaisedInLayer()
	if isRaisedInLayer != retIsRaisedInLayer {
		t.Errorf("the isRaisedInLayer was expected to be %t, %t returned", isRaisedInLayer, retIsRaisedInLayer)
		return
	}
}

func TestFailure_isRaisedInLayer_Success(t *testing.T) {
	index := uint(32)
	code := uint(56)
	isRaisedInLayer := true
	ins := NewFailureForTests(index, code, isRaisedInLayer)
	retCode := ins.Code()
	if code != retCode {
		t.Errorf("the code was expected to be %d, %d returned", code, retCode)
		return
	}

	retIndex := ins.Index()
	if index != retIndex {
		t.Errorf("the index was expected to be %d, %d returned", index, retIndex)
		return
	}

	retIsRaisedInLayer := ins.IsRaisedInLayer()
	if isRaisedInLayer != retIsRaisedInLayer {
		t.Errorf("the isRaisedInLayer was expected to be %t, %t returned", isRaisedInLayer, retIsRaisedInLayer)
		return
	}
}

func TestFailure_withoutCode_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithIndex(uint(54)).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestFailure_withoutIndex_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithCode(uint(22)).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
