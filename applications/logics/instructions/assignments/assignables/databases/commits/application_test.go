package commits

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {

	action := actions.NewActionWithModificationsForTests(
		[]string{"this", "is", "a", "path"},
		modifications.NewModificationsForTests([]modifications.Modification{
			modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
			modifications.NewModificationWithDeleteForTests(
				deletes.NewDeleteForTests(uint(23), uint(56)),
			),
		}),
	)

	actionsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithActionForTests(
			action,
		),
	})

	description := "this is a description"

	actionsVar := "myActions"
	descriptionVar := "myDescription"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				actionsVar,
				stacks.NewAssignableWithListForTests(
					actionsList,
				),
			),
			stacks.NewAssignmentForTests(
				descriptionVar,
				stacks.NewAssignableWithStringForTests(
					description,
				),
			),
		}),
	)

	instruction := commits.NewCommitForTests(descriptionVar, actionsVar)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsCommit() {
		t.Errorf("the assignable was expected to contain a commit")
		return
	}

	retCommit := retAssignable.Commit()
	retActionsList := retCommit.Content().Actions().List()
	if len(retActionsList) != 1 {
		t.Errorf("the returned actions is invalid")
		return
	}

	if !reflect.DeepEqual(action, retActionsList[0]) {
		t.Errorf("the returned action is invalid")
		return
	}

	if description != retCommit.Content().Description() {
		t.Errorf("the returned description is invalid")
		return
	}

	if retCommit.HasParent() {
		t.Errorf("the returned commit was expected to NOT contain a parent hash")
		return
	}
}

func TestExecute_withParent_Success(t *testing.T) {

	action := actions.NewActionWithModificationsForTests(
		[]string{"this", "is", "a", "path"},
		modifications.NewModificationsForTests([]modifications.Modification{
			modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
			modifications.NewModificationWithDeleteForTests(
				deletes.NewDeleteForTests(uint(23), uint(56)),
			),
		}),
	)

	actionsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithActionForTests(
			action,
		),
	})

	description := "this is a description"

	pParent, err := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	if err != nil {
		t.Errorf("the parent was expected to be nil, error returned: %s", err.Error())
		return
	}

	actionsVar := "myActions"
	descriptionVar := "myDescription"
	parentVar := "myParentHash"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				actionsVar,
				stacks.NewAssignableWithListForTests(
					actionsList,
				),
			),
			stacks.NewAssignmentForTests(
				descriptionVar,
				stacks.NewAssignableWithStringForTests(
					description,
				),
			),
			stacks.NewAssignmentForTests(
				parentVar,
				stacks.NewAssignableWithHashForTests(
					*pParent,
				),
			),
		}),
	)

	instruction := commits.NewCommitWithParentForTests(descriptionVar, actionsVar, parentVar)
	application := NewApplication()
	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsCommit() {
		t.Errorf("the assignable was expected to contain a commit")
		return
	}

	retCommit := retAssignable.Commit()
	retActionsList := retCommit.Content().Actions().List()
	if len(retActionsList) != 1 {
		t.Errorf("the returned actions is invalid")
		return
	}

	if !reflect.DeepEqual(action, retActionsList[0]) {
		t.Errorf("the returned action is invalid")
		return
	}

	if description != retCommit.Content().Description() {
		t.Errorf("the returned description is invalid")
		return
	}

	if !retCommit.HasParent() {
		t.Errorf("the returned commit was expected to contain a parent hash")
		return
	}

	if !pParent.Compare(retCommit.Parent()) {
		t.Errorf("the returned parent is invalid")
		return
	}
}

func TestExecute_actionsNotInFrame_returnsError(t *testing.T) {
	description := "this is a description"

	actionsVar := "myActions"
	descriptionVar := "myDescription"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				descriptionVar,
				stacks.NewAssignableWithStringForTests(
					description,
				),
			),
		}),
	)

	instruction := commits.NewCommitForTests(descriptionVar, actionsVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchListFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchListFromFrame)
		return
	}
}

func TestExecute_invalidElementInActions_returnsError(t *testing.T) {
	actionsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithStringForTests(
			"this is an invalid element",
		),
	})

	description := "this is a description"

	actionsVar := "myActions"
	descriptionVar := "myDescription"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				actionsVar,
				stacks.NewAssignableWithListForTests(
					actionsList,
				),
			),
			stacks.NewAssignmentForTests(
				descriptionVar,
				stacks.NewAssignableWithStringForTests(
					description,
				),
			),
		}),
	)

	instruction := commits.NewCommitForTests(descriptionVar, actionsVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchActionFromList {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchActionFromList)
		return
	}
}

func TestExecute_descriptionNotInFrame_returnsError(t *testing.T) {

	action := actions.NewActionWithModificationsForTests(
		[]string{"this", "is", "a", "path"},
		modifications.NewModificationsForTests([]modifications.Modification{
			modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
			modifications.NewModificationWithDeleteForTests(
				deletes.NewDeleteForTests(uint(23), uint(56)),
			),
		}),
	)

	actionsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithActionForTests(
			action,
		),
	})

	actionsVar := "myActions"
	descriptionVar := "myDescription"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				actionsVar,
				stacks.NewAssignableWithListForTests(
					actionsList,
				),
			),
		}),
	)

	instruction := commits.NewCommitForTests(descriptionVar, actionsVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchStringFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchStringFromFrame)
		return
	}
}
