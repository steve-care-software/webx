package histories

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	history := NewHistoryForTests(19)
	adapter := NewAdapter()
	content, err := adapter.ToContent(history)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retHistory, err := adapter.ToHistory(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(history, retHistory) {
		t.Errorf("the returned history is invalid")
		return
	}
}
