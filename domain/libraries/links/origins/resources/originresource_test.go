package resources

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
)

func TestResource_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	originResource := NewResourceForTests(*pLayer)
	retLayer := originResource.Layer()
	if !bytes.Equal(pLayer.Bytes(), retLayer.Bytes()) {
		t.Errorf("the returned layer is invalid")
		return
	}

	if originResource.IsMandatory() {
		t.Errorf("the originResource was expected to NOT be mandatory")
		return
	}

}

func TestResource_isMandatory_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	originResource := NewResourceWithIsMandatoryForTests(*pLayer)
	retLayer := originResource.Layer()
	if !bytes.Equal(pLayer.Bytes(), retLayer.Bytes()) {
		t.Errorf("the returned layer is invalid")
		return
	}

	if !originResource.IsMandatory() {
		t.Errorf("the originResource was expected to be mandatory")
		return
	}

}

func TestResource_withoutLayer_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().IsMandatory().Now()
	if err == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
		return
	}

}
