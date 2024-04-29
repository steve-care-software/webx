package databases

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/reverts"
)

func TestAdapter_withInsert_Success(t *testing.T) {
	ins := databases.NewDatabaseWithInsertForTests(
		inserts.NewInsertForTests(
			"myContext",
			"myInstance",
			"myPath",
		),
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

func TestAdapter_withDelete_Success(t *testing.T) {
	ins := databases.NewDatabaseWithDeleteForTests(
		deletes.NewDeleteForTests(
			"myContext",
			"myPath",
			"myIdentifier",
		),
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

func TestAdapter_withCommit_Success(t *testing.T) {
	ins := databases.NewDatabaseWithCommitForTests(
		"myCommit",
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

func TestAdapter_withCancel_Success(t *testing.T) {
	ins := databases.NewDatabaseWithCancelForTests(
		"myCancel",
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

func TestAdapter_withRevert_Success(t *testing.T) {
	ins := databases.NewDatabaseWithRevertForTests(
		reverts.NewRevertWithIndexForTests(
			"myIndex",
		),
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
