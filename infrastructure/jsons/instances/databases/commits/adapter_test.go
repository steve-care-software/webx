package commits

import (
	"bytes"
	"testing"
	"time"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
)

func TestAdapter_Success(t *testing.T) {

	ins := commits.NewCommitForTests(
		"This is a description",
		actions.NewActionsForTests([]actions.Action{
			actions.NewActionWithModificationsForTests(
				[]string{"this", "is", "a", "path"},
				modifications.NewModificationsForTests([]modifications.Modification{
					modifications.NewModificationWithInsertForTests([]byte("some data to insert")),
					modifications.NewModificationWithDeleteForTests(
						deletes.NewDeleteForTests(
							0,
							50,
						),
					),
				}),
			),
		}),
		time.Now().UTC(),
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

func TestAdapter_withParent_Success(t *testing.T) {

	pHash, err := hash.NewAdapter().FromBytes([]byte("lets say this is bytes"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := commits.NewCommitWithParentForTests(
		"This is the root description",
		actions.NewActionsForTests([]actions.Action{
			actions.NewActionWithModificationsForTests(
				[]string{"this", "is", "a", "path"},
				modifications.NewModificationsForTests([]modifications.Modification{
					modifications.NewModificationWithInsertForTests([]byte("some data to insert")),
					modifications.NewModificationWithDeleteForTests(
						deletes.NewDeleteForTests(
							0,
							50,
						),
					),
				}),
			),
		}),
		time.Now().UTC(),
		*pHash,
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
