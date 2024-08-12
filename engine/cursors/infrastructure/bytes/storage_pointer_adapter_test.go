package bytes

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
)

func TestPointerAdapter_single_Success(t *testing.T) {
	pointer := storages.NewStorageForTests(
		delimiters.NewDelimiterForTests(0, 12),
		true,
	)

	adapter := NewStoragePointerAdapter()
	retBytes, err := adapter.InstanceToBytes(pointer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(pointer, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestPointerAdapter_multiple_Success(t *testing.T) {
	pointers := storages.NewStoragesForTests([]storages.Storage{
		storages.NewStorageForTests(
			delimiters.NewDelimiterForTests(0, 12),
			true,
		),
		storages.NewStorageForTests(
			delimiters.NewDelimiterForTests(1, 22),
			false,
		),
	})

	adapter := NewStoragePointerAdapter()
	retBytes, err := adapter.InstancesToBytes(pointers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(pointers, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
