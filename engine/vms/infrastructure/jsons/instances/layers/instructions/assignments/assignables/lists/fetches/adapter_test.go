package fetches

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
)

func TestAdapter_Success(t *testing.T) {
	ins := fetches.NewFetchForTests(
		"myList",
		"myIndex",
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
