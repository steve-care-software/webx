package assignables

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/bytes"
)

func TestAssignable_withBytes_Success(t *testing.T) {
	bytes := bytes.NewBytesWithJoinForTests([]string{
		"first",
		"second",
	})

	ins := NewAssignableWithBytesForTests(bytes)

	if !ins.IsBytes() {
		t.Errorf("the bytes was expected to contain a bytes")
		return
	}

	retBytes := ins.Bytes()
	if !reflect.DeepEqual(bytes, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}

func TestAssignable_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}
