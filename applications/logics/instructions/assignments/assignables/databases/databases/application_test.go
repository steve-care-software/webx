package databases

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
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

	instruction := databases.NewDatabaseForTests(pathVar, descriptionVar, commitVar, isActiveVar)

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

	if !retAssignable.IsDatabase() {
		t.Errorf("the assignable was expected to contain a database")
		return
	}
}

func TestExecute_pathNotInFrame_returnsError(t *testing.T) {
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

	instruction := databases.NewDatabaseForTests(pathVar, descriptionVar, commitVar, isActiveVar)

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

func TestExecute_descriptionNotInFrame_returnsError(t *testing.T) {
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

	instruction := databases.NewDatabaseForTests(pathVar, descriptionVar, commitVar, isActiveVar)

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

func TestExecute_commitNotInFrame_returnsError(t *testing.T) {
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
				isActiveVar,
				stacks.NewAssignableWithBoolForTests(
					isActive,
				),
			),
		}),
	)

	instruction := databases.NewDatabaseForTests(pathVar, descriptionVar, commitVar, isActiveVar)

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

	if *pCode != failures.CouldNotFetchCommitFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchCommitFromFrame)
		return
	}
}

func TestExecute_isActiveNotInFrame_returnsError(t *testing.T) {
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
		}),
	)

	instruction := databases.NewDatabaseForTests(pathVar, descriptionVar, commitVar, isActiveVar)

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

	if *pCode != failures.CouldNotFetchBoolFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchBoolFromFrame)
		return
	}
}

func TestExecute_withInvalidPathElement_returnsError(t *testing.T) {
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
		stacks.NewAssignableWithBytesForTests(
			[]byte("invalid"),
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

	instruction := databases.NewDatabaseForTests(pathVar, descriptionVar, commitVar, isActiveVar)

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

	if *pCode != failures.CouldNotFetchStringFromList {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchStringFromList)
		return
	}
}
