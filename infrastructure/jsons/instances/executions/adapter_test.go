package executions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
	source_links "github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions/resources"
)

func TestAdapter_Success(t *testing.T) {
	ins := executions.NewExecutionsForTests([]executions.Execution{
		executions.NewExecutionForTests(
			links.NewLinkForTests(
				[]byte("this is some link input"),
				source_links.NewLinkForTests(
					elements.NewElementsForTests([]elements.Element{
						elements.NewElementWithConditionForTests(
							[]string{"path", "to", "layer"},
							conditions.NewConditionForTests(
								resources.NewResourceForTests(uint(45)),
							),
						),
						elements.NewElementForTests(
							[]string{"another", "path", "to", "layer"},
						),
					}),
				),
			),
			databases.NewDatabaseForTests(
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
			),
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
