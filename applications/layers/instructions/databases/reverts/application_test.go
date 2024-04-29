package reverts

import (
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_revertSucceeds_Success(t *testing.T) {
	instruction := reverts.NewRevertForTests()
	frame := stacks.NewFrameForTests()

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(service)
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

func TestExecute_revertFails_returnsError(t *testing.T) {
	instruction := reverts.NewRevertForTests()
	frame := stacks.NewFrameForTests()

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, true)

	application := NewApplication(service)
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
	if code != failures.CouldNotRevertInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotRevertInDatabase, code)
		return
	}
}

func TestExecute_withIndex_indexExistsInFrame_revertSucceeds_Success(t *testing.T) {
	indexVar := "myInstance"
	index := uint(45)
	instruction := reverts.NewRevertWithIndexForTests(indexVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(index),
			),
		}),
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(service)
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

func TestExecute_withIndex_indexExistsInFrame_revertFails_returnsError(t *testing.T) {
	indexVar := "myInstance"
	index := uint(45)
	instruction := reverts.NewRevertWithIndexForTests(indexVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(index),
			),
		}),
	)

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, true)

	application := NewApplication(service)
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
	if code != failures.CouldNotRevertToIndexInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotRevertToIndexInDatabase, code)
		return
	}
}

func TestExecute_withIndex_indexDoesNotExistsInFrame_returnsError(t *testing.T) {
	indexVar := "myInstance"
	instruction := reverts.NewRevertWithIndexForTests(indexVar)
	frame := stacks.NewFrameForTests()

	contextValue := uint(45)
	service := mocks.NewInstanceService(&contextValue, false)

	application := NewApplication(service)
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
	if code != failures.CouldNotFetchIndexFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchIndexFromFrame, code)
		return
	}
}
