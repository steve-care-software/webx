package databases

import (
	"testing"

	application_deletes "github.com/steve-care-software/datastencil/applications/links/layers/instructions/databases/deletes"
	application_inserts "github.com/steve-care-software/datastencil/applications/links/layers/instructions/databases/inserts"
	application_reverts "github.com/steve-care-software/datastencil/applications/links/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_WithInsert_Success(t *testing.T) {
	contextVar := "myContext"
	context := uint(54)
	pathVar := "myPath"
	path := []string{
		"this",
		"is",
		"a",
		"path",
	}

	instanceVar := "myInstance"
	instance := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)

	instruction := databases.NewDatabaseWithInsertForTests(
		inserts.NewInsertForTests(contextVar, instanceVar, pathVar),
	)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				contextVar,
				stacks.NewAssignableWithUnsignedIntForTests(context),
			),
			stacks.NewAssignmentForTests(
				instanceVar,
				stacks.NewAssignableWithInstanceForTests(instance),
			),
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithStringListForTests(path),
			),
		}),
	)

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_WithDelete_Success(t *testing.T) {
	contextVar := "myContext"
	context := uint(54)
	pathVar := "myPath"
	path := []string{
		"this",
		"is",
		"a",
		"path",
	}

	identifierVar := "myHash"
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instance := databases.NewDatabaseWithDeleteForTests(
		deletes.NewDeleteForTests(contextVar, pathVar, identifierVar),
	)

	instruction := databases.NewDatabaseWithDeleteForTests(
		deletes.NewDeleteForTests(contextVar, pathVar, identifierVar),
	)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				contextVar,
				stacks.NewAssignableWithUnsignedIntForTests(context),
			),
			stacks.NewAssignmentForTests(
				pathVar,
				stacks.NewAssignableWithStringListForTests(path),
			),
			stacks.NewAssignmentForTests(
				identifierVar,
				stacks.NewAssignableWithHashForTests(*pHash),
			),
		}),
	)

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{
			pHash.String(): instance,
		},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_WithRevert_Success(t *testing.T) {
	instruction := databases.NewDatabaseWithRevertForTests(
		reverts.NewRevertForTests(),
	)

	frame := stacks.NewFrameForTests()

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_WithCommit_commitExistsInFrame_commitSucceeds_Success(t *testing.T) {
	commitVar := "myContext"
	context := uint(54)

	instruction := databases.NewDatabaseWithCommitForTests(commitVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				commitVar,
				stacks.NewAssignableWithUnsignedIntForTests(context),
			),
		}),
	)

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_WithCommit_commitExistsInFrame_commitFails_returnsError(t *testing.T) {
	commitVar := "myContext"
	context := uint(54)

	instruction := databases.NewDatabaseWithCommitForTests(commitVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				commitVar,
				stacks.NewAssignableWithUnsignedIntForTests(context),
			),
		}),
	)

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, true)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotCommitInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotCommitInDatabase, code)
		return
	}
}

func TestExecute_WithCommit_commitDoesNotExistsInFrame_returnsError(t *testing.T) {
	commitVar := "myContext"

	instruction := databases.NewDatabaseWithCommitForTests(commitVar)
	frame := stacks.NewFrameForTests()

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchContextFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchContextFromFrame, code)
		return
	}
}

func TestExecute_WithCancel_commitExistsInFrame_cancelSucceeds_Success(t *testing.T) {
	cancelVar := "myContext"
	context := uint(54)

	instruction := databases.NewDatabaseWithCancelForTests(cancelVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cancelVar,
				stacks.NewAssignableWithUnsignedIntForTests(context),
			),
		}),
	)

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_WithCancel_cancelExistsInFrame_commitFails_returnsError(t *testing.T) {
	cancelVar := "myContext"
	context := uint(54)

	instruction := databases.NewDatabaseWithCancelForTests(cancelVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cancelVar,
				stacks.NewAssignableWithUnsignedIntForTests(context),
			),
		}),
	)

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, true)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotCancelInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotCancelInDatabase, code)
		return
	}
}

func TestExecute_WithCancel_cancelDoesNotExistsInFrame_returnsError(t *testing.T) {
	cancelVar := "myContext"

	instruction := databases.NewDatabaseWithCancelForTests(cancelVar)
	frame := stacks.NewFrameForTests()

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(
		application_deletes.NewApplication(repository, service),
		application_inserts.NewApplication(repository, service),
		application_reverts.NewApplication(service),
		service,
	)

	pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchContextFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchContextFromFrame, code)
		return
	}
}
