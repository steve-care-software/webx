package deletes

import (
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_contextExistsInFrame_pathExistsInFrame_identifierExistsInFrame_instanceExistsInDatabase_deleteSucceeds_Success(t *testing.T) {
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

	instruction := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)
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

	instance := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)

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

	application := NewApplication(repository, service)
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

func TestExecute_contextExistsInFrame_pathExistsInFrame_identifierExistsInFrame_instanceExistsInDatabase_deleteFails_(t *testing.T) {
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

	instruction := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)
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

	instance := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)

	height := uint(33)
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{
			pHash.String(): instance,
		},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, true)

	application := NewApplication(repository, service)
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
	if code != failures.CouldNotDeleteFromDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotDeleteFromDatabase, code)
		return
	}
}

func TestExecute_contextExistsInFrame_pathExistsInFrame_identifierExistsInFrame_instanceDoesNotExistsInDatabase_returnsError(t *testing.T) {
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

	instruction := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)
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
		map[string]instances.Instance{},
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(repository, service)
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
	if code != failures.InstanceDoesNotExistsInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.InstanceDoesNotExistsInDatabase, code)
		return
	}
}

func TestExecute_contextExistsInFrame_pathExistsInFrame_identifierDoesNotExistsInFrame_returnsError(t *testing.T) {
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

	instruction := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)
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
		}),
	)

	instance := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)

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

	application := NewApplication(repository, service)
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
	if code != failures.CouldNotFetchIdentifierFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchIdentifierFromFrame, code)
		return
	}
}

func TestExecute_contextExistsInFrame_pathDoesNotExistsInFrame_returnsError(t *testing.T) {
	contextVar := "myContext"
	context := uint(54)
	pathVar := "myPath"

	identifierVar := "myHash"
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instruction := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				contextVar,
				stacks.NewAssignableWithUnsignedIntForTests(context),
			),
			stacks.NewAssignmentForTests(
				identifierVar,
				stacks.NewAssignableWithHashForTests(*pHash),
			),
		}),
	)

	instance := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)

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

	application := NewApplication(repository, service)
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
	if code != failures.CouldNotFetchPathFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchPathFromFrame, code)
		return
	}
}

func TestExecute_contextDoesNotExistsInFrame_returnsError(t *testing.T) {
	contextVar := "myContext"
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

	instruction := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
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

	instance := deletes.NewDeleteForTests(contextVar, pathVar, identifierVar)

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

	application := NewApplication(repository, service)
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
