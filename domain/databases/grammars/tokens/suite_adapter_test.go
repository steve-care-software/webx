package tokens

import (
	"reflect"
	"testing"
)

func TestSuiteAdapter_isValid_Success(t *testing.T) {
	suite := NewSuiteForTests(true)
	adapter := NewSuiteAdapter()
	content, err := adapter.ToContent(suite)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retSuite, err := adapter.ToSuite(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(suite, retSuite) {
		t.Errorf("the returned suite is invalid")
		return
	}
}

func TestSuiteAdapter_isInvalidValid_Success(t *testing.T) {
	suite := NewSuiteForTests(false)
	adapter := NewSuiteAdapter()
	content, err := adapter.ToContent(suite)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retSuite, err := adapter.ToSuite(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(suite, retSuite) {
		t.Errorf("the returned suite is invalid")
		return
	}
}
