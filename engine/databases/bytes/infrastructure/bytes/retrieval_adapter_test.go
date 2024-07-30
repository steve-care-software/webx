package bytes

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
)

func TestRetrievalAdapter_single_Success(t *testing.T) {
	retrieval := retrievals.NewRetrievalForTests(0, 12)
	adapter := NewRetrievalAdapter()
	retBytes, err := adapter.InstanceToBytes(retrieval)
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

	if !reflect.DeepEqual(retrieval, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestRetrievalAdapter_single_withRemaining_Success(t *testing.T) {
	retrieval := retrievals.NewRetrievalForTests(0, 12)
	adapter := NewRetrievalAdapter()
	retBytes, err := adapter.InstanceToBytes(retrieval)
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

	if !reflect.DeepEqual(retrieval, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestRetrievalAdapter_multiple_Success(t *testing.T) {

	retrievals := retrievals.NewRetrievalsForTests([]retrievals.Retrieval{
		retrievals.NewRetrievalForTests(0, 12),
		retrievals.NewRetrievalForTests(1, 33),
	})

	adapter := NewRetrievalAdapter()
	retBytes, err := adapter.InstancesToBytes(retrievals)
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

	if !reflect.DeepEqual(retrievals, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestRetrievalAdapter_multiple_withRemaining_Success(t *testing.T) {

	retrievals := retrievals.NewRetrievalsForTests([]retrievals.Retrieval{
		retrievals.NewRetrievalForTests(0, 12),
		retrievals.NewRetrievalForTests(1, 33),
	})

	adapter := NewRetrievalAdapter()
	retBytes, err := adapter.InstancesToBytes(retrievals)
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

	if !reflect.DeepEqual(retrievals, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
