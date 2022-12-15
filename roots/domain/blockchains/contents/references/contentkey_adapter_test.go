package references

import (
	"reflect"
	"testing"
)

func TestContentKeyAdapter_Success(t *testing.T) {
	contentKey := NewContentKeyForTests()
	adapter := NewContentKeyAdapter()
	content, err := adapter.ToContent(contentKey)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retContentKey, err := adapter.ToContentKey(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(contentKey, retContentKey) {
		t.Errorf("the returned contentKey is invalid")
		return
	}
}
