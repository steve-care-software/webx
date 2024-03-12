package accounts

import (
	"testing"

	application_inserts "github.com/steve-care-software/datastencil/applications/layers/instructions/accounts/inserts"
	application_updates "github.com/steve-care-software/datastencil/applications/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/updates/criterias"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stack_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_withInsert_Success(t *testing.T) {
	userVar := "username"
	username := "myUsername"
	passVar := "password"
	password := "myPassword"
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
	application := NewApplication(
		application_inserts.NewApplication(repository, service, bitRate),
		application_updates.NewApplication(repository, service, bitRate),
		service,
	)

	instruction := accounts.NewAccountWithInsertForTests(
		inserts.NewInsertForTests(userVar, passVar),
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

func TestExecute_withUpdate_Success(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	originalCredentialsVar := "originalCredentials"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				originalCredentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(username, []byte(password)),
					),
				),
			),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		false,
		true,
		false,
	)

	bitRate := 4096
	application := NewApplication(
		application_inserts.NewApplication(repository, service, bitRate),
		application_updates.NewApplication(repository, service, bitRate),
		service,
	)

	instruction := accounts.NewAccountWithUpdateForTests(
		updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaForTests(
			true,
			false,
		)),
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

func TestExecute_withDelete_credentialsExistsInFrame_deleteSucceeds_Success(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	deleteCredentialsVar := "originalCredentials"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				deleteCredentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(username, []byte(password)),
					),
				),
			),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		false,
		false,
		true,
	)

	bitRate := 4096
	application := NewApplication(
		application_inserts.NewApplication(repository, service, bitRate),
		application_updates.NewApplication(repository, service, bitRate),
		service,
	)

	instruction := accounts.NewAccountWithDeleteForTests(deleteCredentialsVar)
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

func TestExecute_withDelete_credentialsDoesNotExistsInFrame_deleteSucceeds_ReturnsError(t *testing.T) {
	deleteCredentialsVar := "originalCredentials"
	frame := stacks.NewFrameForTests()
	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		false,
		false,
		true,
	)

	bitRate := 4096
	application := NewApplication(
		application_inserts.NewApplication(repository, service, bitRate),
		application_updates.NewApplication(repository, service, bitRate),
		service,
	)

	instruction := accounts.NewAccountWithDeleteForTests(deleteCredentialsVar)
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
	if code != failures.CouldNotFetchCredentialsFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchCredentialsFromFrame, code)
		return
	}
}

func TestExecute_withDelete_credentialsDoesNotExistsInFrame_deleteFails_ReturnsError(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	deleteCredentialsVar := "originalCredentials"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				deleteCredentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(username, []byte(password)),
					),
				),
			),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		false,
		false,
		false,
	)

	bitRate := 4096
	application := NewApplication(
		application_inserts.NewApplication(repository, service, bitRate),
		application_updates.NewApplication(repository, service, bitRate),
		service,
	)

	instruction := accounts.NewAccountWithDeleteForTests(deleteCredentialsVar)
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
	if code != failures.CouldNotDeleteAccountFromDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotDeleteAccountFromDatabase, code)
		return
	}
}
