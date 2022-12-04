package tokens

import (
	"reflect"
	"testing"
)

func TestLineAdapter_withOneElement_Success(t *testing.T) {
	line := NewLineForTests(1)
	adapter := NewLineAdapter()
	content, err := adapter.ToContent(line)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLine, err := adapter.ToLine(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(line, retLine) {
		t.Errorf("the returned line is invalid")
		return
	}
}

func TestLineAdapter_withMultipleElements_Success(t *testing.T) {
	line := NewLineForTests(100)
	adapter := NewLineAdapter()
	content, err := adapter.ToContent(line)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLine, err := adapter.ToLine(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(line, retLine) {
		t.Errorf("the returned line is invalid")
		return
	}
}
