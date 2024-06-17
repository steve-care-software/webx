package references

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/references"
)

func TestAdapter_Success(t *testing.T) {
	ins := references.NewReferencesForTests([]references.Reference{
		references.NewReferenceForTests(
			"myVariable",
			[]string{"this", "is", "a", "path"},
		),
	})

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
