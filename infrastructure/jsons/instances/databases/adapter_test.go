package databases

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
)

func TestAdapter_Success(t *testing.T) {

	ins := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{"this", "is", "a", "path"},
			"This is the database description",
			true,
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
