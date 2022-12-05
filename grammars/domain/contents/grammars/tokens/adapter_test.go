package tokens

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	token := NewTokenForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(token)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retToken, err := adapter.ToToken(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(token, retToken) {
		t.Errorf("the returned token is invalid")
		return
	}
}

func TestAdapter_withSuites_Success(t *testing.T) {
	token := NewTokenWithSuitesForTests()
	adapter := NewAdapter()
	content, err := adapter.ToContent(token)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retToken, err := adapter.ToToken(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(token, retToken) {
		t.Errorf("the returned token is invalid")
		return
	}
}
