package services

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/databases/services"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_withBegin_beginSucceeds_Success(t *testing.T) {
	context := uint(45)
	frame := stacks.NewFrameForTests()
	instruction := services.NewServiceWithBeginForTests()
	service := mocks.NewInstanceService(
		&context,
		false,
	)

	application := NewApplication(
		service,
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

	if !retAssignable.IsUnsignedInt() {
		t.Errorf("the assignable was expected to contain an uint")
		return
	}

	pRetValue := retAssignable.UnsignedInt()
	if !reflect.DeepEqual(context, *pRetValue) {
		t.Errorf("the returned context is invalid")
		return
	}
}

func TestExecute_withBegin_beginFails_returnsError(t *testing.T) {
	frame := stacks.NewFrameForTests()
	instruction := services.NewServiceWithBeginForTests()
	repository := mocks.NewInstanceService(
		nil,
		false,
	)

	application := NewApplication(
		repository,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotBeginTransactionInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotBeginTransactionInDatabase, code)
		return
	}
}
