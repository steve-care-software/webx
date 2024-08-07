package bytes

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/hashes/domain/pointers"
)

func TestPointerAdapter_single_Success(t *testing.T) {
	pointer := pointers.NewPointerForTests(
		delimiters.NewDelimiterForTests(0, 12),
		true,
	)

	adapter := NewPointerAdapter()
	retBytes, err := adapter.InstanceToBytes(pointer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to contain 0 bytes")
		return
	}

	if !reflect.DeepEqual(pointer, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestPointerAdapter_single_withRemaining_Success(t *testing.T) {
	pointer := pointers.NewPointerForTests(
		delimiters.NewDelimiterForTests(0, 12),
		true,
	)

	adapter := NewPointerAdapter()
	retBytes, err := adapter.InstanceToBytes(pointer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some bytes long enough")
	retBytes = append(retBytes, remaining...)
	retIns, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(pointer, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestPointerAdapter_multiple_Success(t *testing.T) {
	pointers := pointers.NewPointersForTests([]pointers.Pointer{
		pointers.NewPointerForTests(
			delimiters.NewDelimiterForTests(0, 12),
			true,
		),
		pointers.NewPointerForTests(
			delimiters.NewDelimiterForTests(1, 22),
			false,
		),
	})

	adapter := NewPointerAdapter()
	retBytes, err := adapter.InstancesToBytes(pointers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, retRemaining, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to contain 0 bytes")
		return
	}

	if !reflect.DeepEqual(pointers, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestPointerAdapter_multiple_withRemaining_Success(t *testing.T) {
	pointers := pointers.NewPointersForTests([]pointers.Pointer{
		pointers.NewPointerForTests(
			delimiters.NewDelimiterForTests(0, 12),
			true,
		),
		pointers.NewPointerForTests(
			delimiters.NewDelimiterForTests(1, 22),
			false,
		),
	})

	adapter := NewPointerAdapter()
	retBytes, err := adapter.InstancesToBytes(pointers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some bytes long enough")
	retBytes = append(retBytes, remaining...)
	retIns, retRemaining, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(pointers, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
