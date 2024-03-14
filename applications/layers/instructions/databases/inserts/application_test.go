package inserts

import (
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_contextExistsInFrame_instanceExistsInFrame_pathExistsInFrame_instanceExistsInDatabase_insertSucceeds_Success(t *testing.T) {
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

	instruction := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)
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

func TestExecute_contextExistsInFrame_instanceExistsInFrame_pathExistsInFrame_instanceExistsInDatabase_insertFails_returnsError(t *testing.T) {
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

	instruction := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)
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
	if code != failures.CouldNotInsertInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotInsertInDatabase, code)
		return
	}
}

func TestExecute_contextExistsInFrame_instanceExistsInFrame_pathExistsInFrame_instanceDoesNotExistsInDatabase_returnsError(t *testing.T) {
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

	instruction := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)
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
		map[string]instances.Instance{
			instance.Hash().String(): instance,
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
	if code != failures.InstanceAlreadyExistsInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.InstanceAlreadyExistsInDatabase, code)
		return
	}
}

func TestExecute_contextExistsInFrame_instanceExistsInFrame_pathDoesNotExistsInFrame_returnsError(t *testing.T) {
	contextVar := "myContext"
	context := uint(54)
	pathVar := "myPath"

	instanceVar := "myInstance"
	instance := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)

	instruction := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)
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
	if code != failures.CouldNotFetchPathFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchPathFromFrame, code)
		return
	}
}

func TestExecute_contextExistsInFrame_instanceDoesNotExistsInFrame_returnsError(t *testing.T) {
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

	instruction := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)
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
	if code != failures.CouldNotFetchInstanceFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchInstanceFromFrame, code)
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

	instanceVar := "myInstance"
	instance := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)

	instruction := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
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
