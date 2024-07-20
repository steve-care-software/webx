package references

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/references"
)

func TestAdapter_withList_Success(t *testing.T) {
	ins := references.NewReferencesForTests([]references.Reference{
		references.NewReferenceForTests(
			"myVariable",
			[]string{"this", "is", "a", "path"},
		),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_Success(t *testing.T) {
	ins := references.NewReferenceForTests(
		"myVariable",
		[]string{"this", "is", "a", "path"},
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
