package links

import "testing"

func TestConditionResource_Success(t *testing.T) {
	code := uint(45)
	conditionResource := NewConditionResourceForTests(code)
	retCode := conditionResource.Code()
	if code != retCode {
		t.Errorf("the code was expecyed to be %d, %d returned", code, retCode)
		return
	}

	if conditionResource.IsRaisedInLayer() {
		t.Errorf("the conditionResource was expected to NOT be raisedInLayer")
		return
	}
}

func TestConditionResource_isRaisedInLayer_Success(t *testing.T) {
	code := uint(45)
	conditionResource := NewConditionResourceWithIsRaisedInLayerForTests(code)
	retCode := conditionResource.Code()
	if code != retCode {
		t.Errorf("the code was expecyed to be %d, %d returned", code, retCode)
		return
	}

	if !conditionResource.IsRaisedInLayer() {
		t.Errorf("the conditionResource was expected to be raisedInLayer")
		return
	}
}

func TestConditionResource_withoutCode_returnsError(t *testing.T) {
	_, err := NewConditionResourceBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
