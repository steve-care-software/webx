package bytes

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/pointers"
)

func TestPointerAdapter_single_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is some data"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pointer := pointers.NewPointerForTests(
		*pHash,
		delimiters.NewDelimiterForTests(0, 12),
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

func TestPointerAdapter_multiple_Success(t *testing.T) {
	pFirstHash, err := hash.NewAdapter().FromBytes([]byte("this is some data"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pSecondHash, err := hash.NewAdapter().FromBytes([]byte("this is some data again"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pointers := pointers.NewPointersForTests([]pointers.Pointer{
		pointers.NewPointerForTests(
			*pFirstHash,
			delimiters.NewDelimiterForTests(0, 12),
		),
		pointers.NewPointerForTests(
			*pSecondHash,
			delimiters.NewDelimiterForTests(12, 33),
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
