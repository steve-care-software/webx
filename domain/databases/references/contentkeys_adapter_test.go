package references

import (
	"reflect"
	"testing"
)

func TestContentKeysAdapter_Success(t *testing.T) {
	list := []ContentKey{
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
	}

	contentKeys, err := NewContentKeysBuilder().Create().WithList(list).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter := NewContentKeysAdapter()
	content, err := adapter.ToContent(contentKeys)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retContentKeys, err := adapter.ToContentKeys(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(contentKeys, retContentKeys) {
		t.Errorf("the returned contentKeys is invalid")
		return
	}
}
