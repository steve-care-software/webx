package databases

import (
	"bytes"
	"reflect"
	"testing"

	application_actions "github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/actions"
	application_commits "github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/commits"
	application_databases "github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/databases"
	application_deletes "github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/deletes"
	application_modifications "github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/modifications"
	application_retrieves "github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/retrieves"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	assignables_databases "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases"
	instructions_actions "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/actions"
	instructions_commits "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/commits"
	instruction_databases "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/databases"
	instruction_deletes "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/deletes"
	instruction_modifications "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/modifications"
	instruction_retrieves "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/mocks"
)

func TestExecute_withAction_Success(t *testing.T) {
	modificationsList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
		),
		stacks.NewAssignableWithModificationForTests(
			modifications.NewModificationWithDeleteForTests(
				deletes.NewDeleteForTests(uint(23), uint(56)),
			),
		),
	})

	pathList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithStringForTests(
			"this",
		),
		stacks.NewAssignableWithStringForTests(
			"is",
		),
		stacks.NewAssignableWithStringForTests(
			"a",
		),
		stacks.NewAssignableWithStringForTests(
			"path",
		),
	})

	modifVar := "myModifications"
	pathVar := "myPath"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				modifVar,
				stacks.NewAssignableWithListForTests(
					modificationsList,
				),
			),
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithListForTests(
					pathList,
				),
			),
		}),
	)

	instruction := assignables_databases.NewDatabaseWithActionForTests(
		instructions_actions.NewActionForTests(pathVar, modifVar),
	)

	application := NewApplication(
		application_actions.NewApplication(),
		application_commits.NewApplication(),
		application_databases.NewApplication(),
		application_deletes.NewApplication(),
		application_modifications.NewApplication(),
		application_retrieves.NewApplication(
			mocks.NewDatabaseRepository(
				nil,
				nil,
				nil,
			),
		),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsAction() {
		t.Errorf("the assignable was expected to contain an action")
		return
	}

	modifications := retAssignable.Action().Modifications().List()
	if len(modifications) != 2 {
		t.Errorf("the modifications list was expected to contain 2 elements")
		return
	}
}

func TestExecute_withCommit_Success(t *testing.T) {
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

	instruction := assignables_databases.NewDatabaseWithCommitForTests(
		instructions_commits.NewCommitForTests(descriptionVar, actionsVar),
	)

	application := NewApplication(
		application_actions.NewApplication(),
		application_commits.NewApplication(),
		application_databases.NewApplication(),
		application_deletes.NewApplication(),
		application_modifications.NewApplication(),
		application_retrieves.NewApplication(
			mocks.NewDatabaseRepository(
				nil,
				nil,
				nil,
			),
		),
	)

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

func TestExecute_withDatabase_Success(t *testing.T) {
	pathList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithStringForTests(
			"this",
		),
		stacks.NewAssignableWithStringForTests(
			"is",
		),
		stacks.NewAssignableWithStringForTests(
			"a",
		),
		stacks.NewAssignableWithStringForTests(
			"path",
		),
	})

	description := "this is a description"

	commit := commits.NewCommitForTests(
		"this is a commit description",
		actions.NewActionsForTests([]actions.Action{
			actions.NewActionWithModificationsForTests(
				[]string{"this", "is", "a", "path"},
				modifications.NewModificationsForTests([]modifications.Modification{
					modifications.NewModificationWithInsertForTests([]byte("insert some bytes")),
					modifications.NewModificationWithDeleteForTests(
						deletes.NewDeleteForTests(uint(23), uint(56)),
					),
				}),
			),
		}),
	)

	isActive := true

	pathVar := "myPath"
	descriptionVar := "myDescription"
	commitVar := "myCommit"
	isActiveVar := "isActive"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithListForTests(
					pathList,
				),
			),
			stacks.NewAssignmentForTests(
				descriptionVar,
				stacks.NewAssignableWithStringForTests(
					description,
				),
			),
			stacks.NewAssignmentForTests(
				commitVar,
				stacks.NewAssignableWithCommitForTests(
					commit,
				),
			),
			stacks.NewAssignmentForTests(
				isActiveVar,
				stacks.NewAssignableWithBoolForTests(
					isActive,
				),
			),
		}),
	)

	instruction := assignables_databases.NewDatabaseWithDatabaseForTests(
		instruction_databases.NewDatabaseForTests(pathVar, descriptionVar, commitVar, isActiveVar),
	)

	application := NewApplication(
		application_actions.NewApplication(),
		application_commits.NewApplication(),
		application_databases.NewApplication(),
		application_deletes.NewApplication(),
		application_modifications.NewApplication(),
		application_retrieves.NewApplication(
			mocks.NewDatabaseRepository(
				nil,
				nil,
				nil,
			),
		),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsDatabase() {
		t.Errorf("the assignable was expected to contain a database")
		return
	}
}

func TestExecute_withDelete_Success(t *testing.T) {
	index := uint(45)
	length := uint(32)

	indexVar := "myIndex"
	lengthVar := "myLength"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				lengthVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					length,
				),
			),
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					index,
				),
			),
		}),
	)

	instruction := assignables_databases.NewDatabaseWithDeleteForTests(
		instruction_deletes.NewDeleteForTests(indexVar, lengthVar),
	)

	application := NewApplication(
		application_actions.NewApplication(),
		application_commits.NewApplication(),
		application_databases.NewApplication(),
		application_deletes.NewApplication(),
		application_modifications.NewApplication(),
		application_retrieves.NewApplication(
			mocks.NewDatabaseRepository(
				nil,
				nil,
				nil,
			),
		),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsDelete() {
		t.Errorf("the assignable was expected to contain a delete")
		return
	}

	delete := retAssignable.Delete()
	if index != delete.Index() {
		t.Errorf("the index was expected to be %d, %d returned", index, delete.Index())
		return
	}

	if length != delete.Length() {
		t.Errorf("the length was expected to be %d, %d returned", length, delete.Length())
		return
	}
}

func TestExecute_withModification_Success(t *testing.T) {
	insert := []byte("this is some bytes to insert")
	insertVar := "myInsert"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				insertVar,
				stacks.NewAssignableWithBytesForTests(
					insert,
				),
			),
		}),
	)

	instruction := assignables_databases.NewDatabaseWithModificationForTests(
		instruction_modifications.NewModificationWithInsertForTests(insertVar),
	)

	application := NewApplication(
		application_actions.NewApplication(),
		application_commits.NewApplication(),
		application_databases.NewApplication(),
		application_deletes.NewApplication(),
		application_modifications.NewApplication(),
		application_retrieves.NewApplication(
			mocks.NewDatabaseRepository(
				nil,
				nil,
				nil,
			),
		),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsModification() {
		t.Errorf("the assignable was expected to contain a modification")
		return
	}

	modification := retAssignable.Modification()
	if !modification.IsInsert() {
		t.Errorf("the assignable was expected to contain an insert modification")
		return
	}

	retInsert := modification.Insert()
	if !bytes.Equal(insert, retInsert) {
		t.Errorf("the returned insert is invalid")
		return
	}
}

func TestExecute_withRetrieve_Success(t *testing.T) {
	pathList := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithStringForTests(
			"this",
		),
		stacks.NewAssignableWithStringForTests(
			"is",
		),
		stacks.NewAssignableWithStringForTests(
			"a",
		),
		stacks.NewAssignableWithStringForTests(
			"path",
		),
	})

	pathVar := "myPath"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithListForTests(
					pathList,
				),
			),
		}),
	)

	instruction := assignables_databases.NewDatabaseWithRetrieveForTests(
		instruction_retrieves.NewRetrieveWithExistsForTests(pathVar),
	)

	application := NewApplication(
		application_actions.NewApplication(),
		application_commits.NewApplication(),
		application_databases.NewApplication(),
		application_deletes.NewApplication(),
		application_modifications.NewApplication(),
		application_retrieves.NewApplication(
			mocks.NewDatabaseRepository(
				nil,
				nil,
				nil,
			),
		),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain a bool")
		return
	}

	retBool := retAssignable.Bool()
	if *retBool {
		t.Errorf("the returned bool was expected to be false, true returned")
		return
	}
}
