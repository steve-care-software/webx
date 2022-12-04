package tokens

import (
	"reflect"
	"testing"
)

func TestSuitesAdapter_withOneElement_Success(t *testing.T) {
	suites := NewSuitesForTests(1)
	adapter := NewSuitesAdapter()
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

func TestSuitesAdapter_withMultipleElements_Success(t *testing.T) {
	suites := NewSuitesForTests(100)
	adapter := NewSuitesAdapter()
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
