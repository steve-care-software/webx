package contents

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	contents := NewContentsForTests(0, [][]byte{
		[]byte("this is the first data"),
		[]byte("this is the second data"),
		[]byte("this is the third data"),
	})

	adapter := NewAdapter()
	content, err := adapter.ToContent(contents)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, err := adapter.ToInstance(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(contents, retInstance) {
		t.Errorf("the returned contents instance is invalid")
		return
	}
}
