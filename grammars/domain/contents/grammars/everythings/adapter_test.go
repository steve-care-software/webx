package everythings

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	everyting := NewEverythingForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(everyting)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retEverything, err := adapter.ToEverything(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(everyting, retEverything) {
		t.Errorf("the returned everyting is invalid")
		return
	}
}

func TestAdapter_withEscape_Success(t *testing.T) {
	everyting := NewEverythingWithEscapeForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(everyting)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retEverything, err := adapter.ToEverything(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(everyting, retEverything) {
		t.Errorf("the returned everyting is invalid")
		return
	}
}
