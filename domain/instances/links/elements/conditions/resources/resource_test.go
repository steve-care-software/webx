package resources

import "testing"

func TestResource_Success(t *testing.T) {
	code := uint(45)
	resource := NewResourceForTests(code)
	retCode := resource.Code()
	if code != retCode {
		t.Errorf("the code was expecyed to be %d, %d returned", code, retCode)
		return
	}

	if resource.IsRaisedInLayer() {
		t.Errorf("the resource was expected to NOT be raisedInLayer")
		return
	}
}

func TestResource_isRaisedInLayer_Success(t *testing.T) {
	code := uint(45)
	resource := NewResourceWithIsRaisedInLayerForTests(code)
	retCode := resource.Code()
	if code != retCode {
		t.Errorf("the code was expecyed to be %d, %d returned", code, retCode)
		return
	}

	if !resource.IsRaisedInLayer() {
		t.Errorf("the resource was expected to be raisedInLayer")
		return
	}
}

func TestResource_withoutCode_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
