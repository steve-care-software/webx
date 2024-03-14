package inserts

import (
	"testing"

	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_userExistsInFrame_passwordExistsInFrame_userDoesNotExists_insertSucceeds_Success(t *testing.T) {
	userVar := "username"
	username := "myUsername"
	passVar := "password"
	password := "myPassword"
	insert := inserts.NewInsertForTests(userVar, passVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(userVar, stacks.NewAssignableWithBytesForTests([]byte(username))),
			stacks.NewAssignmentForTests(passVar, stacks.NewAssignableWithBytesForTests([]byte(password))),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		true,
		false,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, insert)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_userDoesNotExistsInFrame_passwordExistsInFrame_userDoesNotExists_insertSucceeds_ReturnsError(t *testing.T) {
	userVar := "username"
	username := "myUsername"
	passVar := "password"
	password := "myPassword"
	insert := inserts.NewInsertForTests("nonMatchingUserVar", passVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(userVar, stacks.NewAssignableWithBytesForTests([]byte(username))),
			stacks.NewAssignmentForTests(passVar, stacks.NewAssignableWithBytesForTests([]byte(password))),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		true,
		false,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, insert)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchUsernameFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchUsernameFromFrame, code)
		return
	}
}

func TestExecute_userExistsInFrame_passwordDoesNotExistsInFrame_userDoesNotExists_insertSucceeds_ReturnsError(t *testing.T) {
	userVar := "username"
	username := "myUsername"
	passVar := "password"
	password := "myPassword"
	insert := inserts.NewInsertForTests(userVar, "nonMatchingPassVar")
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(userVar, stacks.NewAssignableWithBytesForTests([]byte(username))),
			stacks.NewAssignmentForTests(passVar, stacks.NewAssignableWithBytesForTests([]byte(password))),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		true,
		false,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, insert)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchPasswordFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchPasswordFromFrame, code)
		return
	}
}

func TestExecute_userExistsInFrame_passwordExistsInFrame_userAlreadyExists_insertSucceeds_Success(t *testing.T) {
	userVar := "username"
	username := "myUsername"
	passVar := "password"
	password := "myPassword"
	insert := inserts.NewInsertForTests(userVar, passVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(userVar, stacks.NewAssignableWithBytesForTests([]byte(username))),
			stacks.NewAssignmentForTests(passVar, stacks.NewAssignableWithBytesForTests([]byte(password))),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{
		username: true,
	})
	service := mocks.NewAccountServiceForTests(
		true,
		false,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, insert)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.AccountWithSameUsernameAlreadyExists {
		t.Errorf("the code was expected to be %d, %d returned", failures.AccountWithSameUsernameAlreadyExists, code)
		return
	}
}

func TestExecute_userExistsInFrame_passwordExistsInFrame_userDoesNotExists_insertFails_ReturnsError(t *testing.T) {
	userVar := "username"
	username := "myUsername"
	passVar := "password"
	password := "myPassword"
	insert := inserts.NewInsertForTests(userVar, passVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(userVar, stacks.NewAssignableWithBytesForTests([]byte(username))),
			stacks.NewAssignmentForTests(passVar, stacks.NewAssignableWithBytesForTests([]byte(password))),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		false,
		false,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, insert)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotInsertAccountInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotInsertAccountInDatabase, code)
		return
	}
}
