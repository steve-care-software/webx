package databases

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
	database_instruction "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
	"github.com/steve-care-software/datastencil/domain/stacks/mocks"
)

func TestExecute_withSave_Success(t *testing.T) {
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

	saveVar := "mySave"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				saveVar,
				stacks.NewAssignableWithDatabaseForTests(
					database,
				),
			),
		}),
	)

	instruction := database_instruction.NewDatabaseWithSaveForTests(saveVar)

	service := mocks.NewDatabaseService(
		database,
	)

	application := NewApplication(
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

}

func TestExecute_withSave_saveReturnsError_returnsError(t *testing.T) {
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

	saveVar := "mySave"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				saveVar,
				stacks.NewAssignableWithDatabaseForTests(
					database,
				),
			),
		}),
	)

	instruction := database_instruction.NewDatabaseWithSaveForTests(saveVar)

	service := mocks.NewDatabaseService(
		databases.NewDatabaseForTests(
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
					"invalid",
					"path",
				},
				"this is a description",
				false,
			),
		),
	)

	application := NewApplication(
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to NOT be nil")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to NOT be nil")
		return
	}

	if *pCode != failures.CouldNotSaveDatabaseFromService {
		t.Errorf("the error code was expected to be %d, %d returned", failures.CouldNotSaveDatabaseFromService, *pCode)
		return
	}

}

func TestExecute_withSave_databaseNotInFrame_returnsError(t *testing.T) {
	saveVar := "mySave"
	frame := stacks.NewFrameForTests()
	instruction := database_instruction.NewDatabaseWithSaveForTests(saveVar)

	service := mocks.NewDatabaseService(
		databases.NewDatabaseForTests(
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
					"invalid",
					"path",
				},
				"this is a description",
				false,
			),
		),
	)

	application := NewApplication(
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to NOT be nil")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to NOT be nil")
		return
	}

	if *pCode != failures.CouldNotFetchDatabaseFromFrame {
		t.Errorf("the error code was expected to be %d, %d returned", failures.CouldNotFetchDatabaseFromFrame, *pCode)
		return
	}

}

func TestExecute_withDelete_Success(t *testing.T) {
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

	deleteVar := "myDelete"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				deleteVar,
				stacks.NewAssignableWithHashForTests(
					database.Hash(),
				),
			),
		}),
	)

	instruction := database_instruction.NewDatabaseWithDeleteForTests(deleteVar)

	service := mocks.NewDatabaseService(
		database,
	)

	application := NewApplication(
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

}

func TestExecute_withDelete_databaseDoesNotExists_returnsError(t *testing.T) {
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

	pHash, err := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	deleteVar := "myDelete"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				deleteVar,
				stacks.NewAssignableWithHashForTests(
					*pHash,
				),
			),
		}),
	)

	instruction := database_instruction.NewDatabaseWithDeleteForTests(deleteVar)

	service := mocks.NewDatabaseService(
		database,
	)

	application := NewApplication(
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to NOT be nil")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to NOT be nil")
		return
	}

	if *pCode != failures.CouldNotDeleteDatabaseFromService {
		t.Errorf("the error code was expected to be %d, %d returned", failures.CouldNotDeleteDatabaseFromService, *pCode)
		return
	}

}

func TestExecute_withDelete_databaseHashIsNotInFrame_returnsError(t *testing.T) {
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

	deleteVar := "myDelete"
	frame := stacks.NewFrameForTests()
	instruction := database_instruction.NewDatabaseWithDeleteForTests(deleteVar)

	service := mocks.NewDatabaseService(
		database,
	)

	application := NewApplication(
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to NOT be nil")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to NOT be nil")
		return
	}

	if *pCode != failures.CouldNotFetchHashVariableFromFrame {
		t.Errorf("the error code was expected to be %d, %d returned", failures.CouldNotFetchHashVariableFromFrame, *pCode)
		return
	}

}
