package contents

import (
	"reflect"
	"testing"
)

func TestContentAdapter_Success(t *testing.T) {
	contentIns := NewContentForTests(0, []byte("this is some data"))
	adapter := NewContentAdapter()
	content, err := adapter.ToContent(contentIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, err := adapter.ToInstance(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(contentIns, retInstance) {
		t.Errorf("the returned content instance is invalid")
		return
	}
}
