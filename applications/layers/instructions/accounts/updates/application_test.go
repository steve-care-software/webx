package updates

import (
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts/updates/criterias"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stack_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_credentialsExistsInFrame_changeSigner_doNotChangeEncryptor_doNotChangeUsername_doNotChangePassword_updateSucceeds_Success(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	originalCredentialsVar := "originalCredentials"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaForTests(
		true,
		false,
	))

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
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_credentialsDoesNotExistsInFrame_changeSigner_doNotChangeEncryptor_doNotChangeUsername_doNotChangePassword_updateSucceeds_Success(t *testing.T) {
	originalCredentialsVar := "originalCredentials"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaForTests(
		true,
		false,
	))

	frame := stacks.NewFrameForTests()
	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	service := mocks.NewAccountServiceForTests(
		false,
		true,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
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

func TestExecute_credentialsExistsInFrame_changeSigner_doNotChangeEncryptor_doNotChangeUsername_doNotChangePassword_updateFails_ReturnsError(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	originalCredentialsVar := "originalCredentials"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaForTests(
		true,
		false,
	))

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
		false,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotUpdateAccountInDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotUpdateAccountInDatabase, code)
		return
	}
}

func TestExecute_credentialsExistsInFrame_doNotChangeSigner_changeEncryptor_doNotChangeUsername_doNotChangePassword_updateSucceeds_Success(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	originalCredentialsVar := "originalCredentials"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaForTests(
		false,
		true,
	))

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
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_credentialsExistsInFrame_doNotChangeSigner_doNotChangeEncryptor_changeUsername_newUsernameExistsInFrame_doNotChangePassword_updateSucceeds_Success(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	newUsername := "newUsername"
	newUsernameVar := "newUsernameVar"
	originalCredentialsVar := "originalCredentialsVar"
	newCredentialsVar := "newCredentialsBar"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaWithUsernameForTests(
		false,
		false,
		newUsernameVar,
	))

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
			stacks.NewAssignmentForTests(
				newCredentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(newUsername, []byte(password)),
					),
				),
			),
			stacks.NewAssignmentForTests(
				newUsernameVar,
				stacks.NewAssignableWithBytesForTests([]byte(newUsername)),
			),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{
		newUsername: false,
	})

	service := mocks.NewAccountServiceForTests(
		false,
		true,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_credentialsExistsInFrame_doNotChangeSigner_doNotChangeEncryptor_changeUsername_newUsernameDoesNotExistsInFrame_doNotChangePassword_updateSucceeds_ReturnsError(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	newUsername := "newUsername"
	newUsernameVar := "newUsernameVar"
	originalCredentialsVar := "originalCredentialsVar"
	newCredentialsVar := "newCredentialsBar"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaWithUsernameForTests(
		false,
		false,
		newUsernameVar,
	))

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
			stacks.NewAssignmentForTests(
				newCredentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(newUsername, []byte(password)),
					),
				),
			),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{
		newUsername: false,
	})

	service := mocks.NewAccountServiceForTests(
		false,
		true,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
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

func TestExecute_credentialsExistsInFrame_doNotChangeSigner_doNotChangeEncryptor_changeUsername_newUsernameExistsInFrame_newUsernameAlreadyExists_doNotChangePassword_updateSucceeds_Success(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	newUsername := "newUsername"
	newUsernameVar := "newUsernameVar"
	originalCredentialsVar := "originalCredentialsVar"
	newCredentialsVar := "newCredentialsBar"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaWithUsernameForTests(
		false,
		false,
		newUsernameVar,
	))

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
			stacks.NewAssignmentForTests(
				newCredentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(newUsername, []byte(password)),
					),
				),
			),
			stacks.NewAssignmentForTests(
				newUsernameVar,
				stacks.NewAssignableWithBytesForTests([]byte(newUsername)),
			),
		}),
	)

	repository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{
		newUsername: true,
	})

	service := mocks.NewAccountServiceForTests(
		false,
		true,
		false,
	)

	bitRate := 4096
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
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

func TestExecute_credentialsExistsInFrame_doNotChangeSigner_doNotChangeEncryptor_doNotChangeUsername_changePassword_passwordExistsInFrame_updateSucceeds_Success(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	newPassword := "thisIsANewPassword"
	newUsername := "newUsername"
	newPasswordVar := "newPasswordVar"
	originalCredentialsVar := "originalCredentialsVar"
	newCredentialsVar := "newCredentialsBar"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaWithPasswordForTests(
		false,
		false,
		newPasswordVar,
	))

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
			stacks.NewAssignmentForTests(
				newCredentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(newUsername, []byte(password)),
					),
				),
			),
			stacks.NewAssignmentForTests(
				newPasswordVar,
				stacks.NewAssignableWithBytesForTests([]byte(newPassword)),
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
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}
}

func TestExecute_credentialsExistsInFrame_doNotChangeSigner_doNotChangeEncryptor_doNotChangeUsername_changePassword_passwordDoesNotExistsInFrame_updateSucceeds_ReturnsError(t *testing.T) {
	username := "myUsername"
	password := "myPassword"
	newUsername := "newUsername"
	newPasswordVar := "newPasswordVar"
	originalCredentialsVar := "originalCredentialsVar"
	newCredentialsVar := "newCredentialsBar"
	update := updates.NewUpdateForTests(originalCredentialsVar, criterias.NewCriteriaWithPasswordForTests(
		false,
		false,
		newPasswordVar,
	))

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
			stacks.NewAssignmentForTests(
				newCredentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(newUsername, []byte(password)),
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
	application := NewApplication(repository, service, bitRate)
	pCode, err := application.Execute(frame, update)
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
