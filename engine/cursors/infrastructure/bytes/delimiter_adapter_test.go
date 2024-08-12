package bytes

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
)

func TestDelimiterAdapter_single_Success(t *testing.T) {
	delimiter := delimiters.NewDelimiterForTests(0, 12)
	adapter := NewDelimiterAdapter()
	retBytes, err := adapter.InstanceToBytes(delimiter)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(delimiter, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestDelimiterAdapter_multiple_Success(t *testing.T) {

	delimiters := delimiters.NewDelimitersForTests([]delimiters.Delimiter{
		delimiters.NewDelimiterForTests(0, 12),
		delimiters.NewDelimiterForTests(1, 33),
	})

	adapter := NewDelimiterAdapter()
	retBytes, err := adapter.InstancesToBytes(delimiters)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(delimiters, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
