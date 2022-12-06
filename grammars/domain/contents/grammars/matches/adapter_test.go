package matches

import (
	"reflect"
	"testing"
)

func TestAdapter_withOneSuite_Success(t *testing.T) {
	match := NewMatchForTests(1)
	adapter := NewAdapter()
	content, err := adapter.ToContent(match)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retMatch, err := adapter.ToMatch(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(match, retMatch) {
		t.Errorf("the returned match is invalid")
		return
	}
}

func TestAdapter_withmultipleSuites_Success(t *testing.T) {
	match := NewMatchForTests(100)
	adapter := NewAdapter()
	content, err := adapter.ToContent(match)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retMatch, err := adapter.ToMatch(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(match, retMatch) {
		t.Errorf("the returned match is invalid")
		return
	}
}
