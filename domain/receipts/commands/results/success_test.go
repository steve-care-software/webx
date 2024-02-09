package results

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/libraries/layers"
)

func TestSuccess_Success(t *testing.T) {
	value := []byte("this is some bytes")
	kind := layers.NewKindWithPromptForTests()
	ins := NewSuccessForTests(value, kind)
	retBytes := ins.Bytes()
	if !bytes.Equal(value, retBytes) {
		t.Errorf("the returned bytes are invalid")
		return
	}

	retKind := ins.Kind()
	if !reflect.DeepEqual(kind, retKind) {
		t.Errorf("the returned kind is invalid")
		return
	}
}

func TestSuccess_withoutBytes_returnsError(t *testing.T) {
	kind := layers.NewKindWithPromptForTests()
	_, err := NewSuccessBuilder().Create().WithKind(kind).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestSuccess_withoutKind_returnsError(t *testing.T) {
	value := []byte("this is some bytes")
	_, err := NewSuccessBuilder().Create().WithBytes(value).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
