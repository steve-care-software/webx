package results

import (
	"testing"
)

func TestFailure_Success(t *testing.T) {
	code := uint(56)
	isRaisedInLayer := false
	ins := NewFailureForTests(code, isRaisedInLayer)
	retCode := ins.Code()
	if code != retCode {
		t.Errorf("the code was expected to be %d, %d returned", code, retCode)
		return
	}

	retIsRaisedInLayer := ins.IsRaisedInLayer()
	if isRaisedInLayer != retIsRaisedInLayer {
		t.Errorf("the isRaisedInLayer was expected to be %t, %t returned", isRaisedInLayer, retIsRaisedInLayer)
		return
	}
}

func TestFailure_isRaisedInLayer_Success(t *testing.T) {
	code := uint(56)
	isRaisedInLayer := true
	ins := NewFailureForTests(code, isRaisedInLayer)
	retCode := ins.Code()
	if code != retCode {
		t.Errorf("the code was expected to be %d, %d returned", code, retCode)
		return
	}

	retIsRaisedInLayer := ins.IsRaisedInLayer()
	if isRaisedInLayer != retIsRaisedInLayer {
		t.Errorf("the isRaisedInLayer was expected to be %t, %t returned", isRaisedInLayer, retIsRaisedInLayer)
		return
	}
}

func TestFailure_withoutCode_returnsError(t *testing.T) {
	_, err := NewFailureBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
