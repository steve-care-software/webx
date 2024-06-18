package databases

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/commits"
	databases_database "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/retrieves"
)

func TestAdapter_withAction_Success(t *testing.T) {
	ins := databases.NewDatabaseWithActionForTests(
		actions.NewActionForTests(
			"myPath",
			"myModifications",
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
		commits.NewCommitForTests(
			"This is a description",
			"myActions",
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

func TestAdapter_withDatabase_Success(t *testing.T) {
	ins := databases.NewDatabaseWithDatabaseForTests(
		databases_database.NewDatabaseForTests(
			"myPath",
			"This is a description",
			"myHead",
			"isActive",
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
			"myIndex",
			"myLength",
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

func TestAdapter_withModification_Success(t *testing.T) {
	ins := databases.NewDatabaseWithModificationForTests(
		modifications.NewModificationWithInsertForTests(
			"myInsert",
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

func TestAdapter_withRetrieve_Success(t *testing.T) {
	ins := databases.NewDatabaseWithRetrieveForTests(
		retrieves.NewRetrieveWithExistsForTests(
			"myExists",
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
