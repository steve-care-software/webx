package references

import (
	"reflect"
	"testing"
)

func TestPointerAdapter_Success(t *testing.T) {
	pointer := NewPointerForTests()
	adapter := NewPointerAdapter()
	content, err := adapter.ToContent(pointer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retPointer, err := adapter.ToPointer(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(pointer, retPointer) {
		t.Errorf("the returned pointer is invalid")
		return
	}
}
