package tokens

import (
	"reflect"
	"testing"
)

func TestLinesAdapter_withOneElement_Success(t *testing.T) {
	lines := NewLinesForTests(1)
	adapter := NewLinesAdapter()
	content, err := adapter.ToContent(lines)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLine, err := adapter.ToLines(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(lines, retLine) {
		t.Errorf("the returned lines is invalid")
		return
	}
}

func TestLinesAdapter_withMultipleElements_Success(t *testing.T) {
	lines := NewLinesForTests(100)
	adapter := NewLinesAdapter()
	content, err := adapter.ToContent(lines)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLine, err := adapter.ToLines(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(lines, retLine) {
		t.Errorf("the returned lines is invalid")
		return
	}
}
