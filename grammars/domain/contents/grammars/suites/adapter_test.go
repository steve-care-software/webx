package suites

import (
	"reflect"
	"testing"
)

func TestAdapter_withOneElement_Success(t *testing.T) {
	suites := NewSuitesForTests(1)
	adapter := NewAdapter()
	content, err := adapter.ToContent(suites)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retSuite, err := adapter.ToSuites(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(suites, retSuite) {
		t.Errorf("the returned suites is invalid")
		return
	}
}

func TestAdapter_withMultipleElements_Success(t *testing.T) {
	suites := NewSuitesForTests(100)
	adapter := NewAdapter()
	content, err := adapter.ToContent(suites)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retSuite, err := adapter.ToSuites(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(suites, retSuite) {
		t.Errorf("the returned suites is invalid")
		return
	}
}
