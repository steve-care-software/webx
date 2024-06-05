package retrieves

import (
	"errors"
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
	"github.com/steve-care-software/datastencil/domain/stacks/mocks"
)

func TestExecute_withList_Success(t *testing.T) {
	pathList := [][]string{
		[]string{
			"this",
			"is",
			"a",
			"path",
			"first",
		},
		[]string{
			"this",
			"is",
			"a",
			"path",
			"second",
		},
	}

	repository := mocks.NewDatabaseRepository(
		pathList,
		nil,
		nil,
	)

	frame := stacks.NewFrameForTests()

	instruction := retrieves.NewRetrieveWithListForTests()
	application := NewApplication(
		repository,
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

	if !retAssignable.IsList() {
		t.Errorf("the assignable was expected to contain a list")
		return
	}

	retList := retAssignable.List().List()
	if len(pathList) != len(retList) {
		t.Errorf("the returned path is invalid")
		return
	}
}

func TestExecute_withList_returnsErrorInRepository_returnsError(t *testing.T) {
	pathList := [][]string{
		[]string{
			"this",
			"is",
			"a",
			"path",
			"first",
		},
		[]string{
			"this",
			"is",
			"a",
			"path",
			"second",
		},
	}

	repository := mocks.NewDatabaseRepository(
		pathList,
		nil,
		errors.New("error"),
	)

	frame := stacks.NewFrameForTests()

	instruction := retrieves.NewRetrieveWithListForTests()
	application := NewApplication(
		repository,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotRetrieveListFromRepository {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotRetrieveListFromRepository)
		return
	}
}

func TestExecute_withList_withEmptyPathList_returnsError(t *testing.T) {
	pathList := [][]string{}
	repository := mocks.NewDatabaseRepository(
		pathList,
		nil,
		nil,
	)

	frame := stacks.NewFrameForTests()

	instruction := retrieves.NewRetrieveWithListForTests()
	application := NewApplication(
		repository,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotRetrieveEmptyListFromRepository {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotRetrieveEmptyListFromRepository)
		return
	}
}

func TestExecute_withDatabase_retrieve_Success(t *testing.T) {

	database := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{
				"this",
				"is",
				"a",
				"path",
			},
			"this is a description",
			false,
		),
	)

	repository := mocks.NewDatabaseRepository(
		nil,
		database,
		nil,
	)

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

	instruction := retrieves.NewRetrieveWithRetrieveForTests(pathVar)
	application := NewApplication(
		repository,
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

	retDatabase := retAssignable.Database()
	if !reflect.DeepEqual(database, retDatabase) {
		t.Errorf("the returned path is invalid")
		return
	}
}

func TestExecute_withDatabase_retrieve_repositoryReturnsError_returnsError(t *testing.T) {

	database := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{
				"this",
				"is",
				"a",
				"path",
			},
			"this is a description",
			false,
		),
	)

	repository := mocks.NewDatabaseRepository(
		nil,
		database,
		errors.New("this is an error"),
	)

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

	instruction := retrieves.NewRetrieveWithRetrieveForTests(pathVar)
	application := NewApplication(
		repository,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotRetrieveFromRepository {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotRetrieveFromRepository)
		return
	}
}

func TestExecute_withDatabase_retrieve_pathContainsInvalidElement_returnsError(t *testing.T) {

	database := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{
				"this",
				"is",
				"a",
				"path",
			},
			"this is a description",
			false,
		),
	)

	repository := mocks.NewDatabaseRepository(
		nil,
		database,
		nil,
	)

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

	instruction := retrieves.NewRetrieveWithRetrieveForTests(pathVar)
	application := NewApplication(
		repository,
	)

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

func TestExecute_withDatabase_retrieve_pathNotInFrame_returnsError(t *testing.T) {

	database := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{
				"this",
				"is",
				"a",
				"path",
			},
			"this is a description",
			false,
		),
	)

	repository := mocks.NewDatabaseRepository(
		nil,
		database,
		nil,
	)

	pathVar := "myPath"

	frame := stacks.NewFrameForTests()

	instruction := retrieves.NewRetrieveWithRetrieveForTests(pathVar)
	application := NewApplication(
		repository,
	)

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

func TestExecute_withDatabase_exists_doesExists_Success(t *testing.T) {

	database := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{
				"this",
				"is",
				"a",
				"path",
			},
			"this is a description",
			false,
		),
	)

	repository := mocks.NewDatabaseRepository(
		nil,
		database,
		nil,
	)

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

	instruction := retrieves.NewRetrieveWithExistsForTests(pathVar)
	application := NewApplication(
		repository,
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
	if !*retBool {
		t.Errorf("the returned bool was expected to be true, false returned")
		return
	}
}

func TestExecute_withDatabase_exists_doesNotExists_Success(t *testing.T) {
	repository := mocks.NewDatabaseRepository(
		nil,
		nil,
		nil,
	)

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

	instruction := retrieves.NewRetrieveWithExistsForTests(pathVar)
	application := NewApplication(
		repository,
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

func TestExecute_withDatabase_exists_repositoryReturnsError_returnsError(t *testing.T) {

	database := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{
				"this",
				"is",
				"a",
				"path",
			},
			"this is a description",
			false,
		),
	)

	repository := mocks.NewDatabaseRepository(
		nil,
		database,
		errors.New("this is an error"),
	)

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

	instruction := retrieves.NewRetrieveWithExistsForTests(pathVar)
	application := NewApplication(
		repository,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotExecuteExistsFromRepository {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotExecuteExistsFromRepository)
		return
	}
}

func TestExecute_withDatabase_exists_invalidElementInPath_returnsError(t *testing.T) {

	database := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{
				"this",
				"is",
				"a",
				"path",
			},
			"this is a description",
			false,
		),
	)

	repository := mocks.NewDatabaseRepository(
		nil,
		database,
		nil,
	)

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

	instruction := retrieves.NewRetrieveWithExistsForTests(pathVar)
	application := NewApplication(
		repository,
	)

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

func TestExecute_withDatabase_exists_pathNotInFrame_returnsError(t *testing.T) {

	database := databases.NewDatabaseForTests(
		commits.NewCommitForTests(
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
		),
		heads.NewHeadForTests(
			[]string{
				"this",
				"is",
				"a",
				"path",
			},
			"this is a description",
			false,
		),
	)

	repository := mocks.NewDatabaseRepository(
		nil,
		database,
		nil,
	)

	pathVar := "myPath"

	frame := stacks.NewFrameForTests()

	instruction := retrieves.NewRetrieveWithExistsForTests(pathVar)
	application := NewApplication(
		repository,
	)

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
