package retrieves

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stack_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_passwordExistsInFrame_credentialsExistsInFrame_accountExistsInDatabase_Success(t *testing.T) {
	username := "myUsername"
	passVar := "password"
	password := "myPassword"

	credentialsVar := "myCredentials"

	instruction := retrieves.NewRetrieveForTests(passVar, credentialsVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				passVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(password),
				),
			),
			stacks.NewAssignmentForTests(
				credentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(username, []byte(password)),
					),
				),
			),
		}),
	)

	account := accounts.NewAccountForTests(username, encryptors.NewEncryptorForTests(4096))
	repository := mocks.NewAccountRepositoryWithRetrieveForTests(map[string]accounts.Account{
		username: account,
	})

	application := NewApplication(repository)
	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsAccount() {
		t.Errorf("the assignable was expected to contain an account")
		return
	}

	retAccount := retAssignable.Account()
	if !retAccount.IsAccount() {
		t.Errorf("the assignable (account) was expected to contain an account")
		return
	}

	retAcc := retAccount.Account()
	if !reflect.DeepEqual(account, retAcc) {
		t.Errorf("the returned account is invalid")
		return
	}
}

func TestExecute_passwordExistsInFrame_credentialsExistsInFrame_accountDoesNotExistsInDatabase_returnsError(t *testing.T) {
	username := "myUsername"
	passVar := "password"
	password := "myPassword"

	credentialsVar := "myCredentials"

	instruction := retrieves.NewRetrieveForTests(passVar, credentialsVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				passVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(password),
				),
			),
			stacks.NewAssignmentForTests(
				credentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(username, []byte(password)),
					),
				),
			),
		}),
	)

	repository := mocks.NewAccountRepositoryWithRetrieveForTests(map[string]accounts.Account{})
	application := NewApplication(repository)
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
	if code != failures.CouldNotRetrieveAccountFromDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotRetrieveAccountFromDatabase, code)
		return
	}
}

func TestExecute_passwordExistsInFrame_credentialsDoesNotExistsInFrame_returnsError(t *testing.T) {
	username := "myUsername"
	passVar := "password"
	password := "myPassword"

	credentialsVar := "myCredentials"

	instruction := retrieves.NewRetrieveForTests(passVar, credentialsVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				passVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(password),
				),
			),
		}),
	)

	account := accounts.NewAccountForTests(username, encryptors.NewEncryptorForTests(4096))
	repository := mocks.NewAccountRepositoryWithRetrieveForTests(map[string]accounts.Account{
		username: account,
	})

	application := NewApplication(repository)
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
	if code != failures.CouldNotFetchCredentialsFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchCredentialsFromFrame, code)
		return
	}
}

func TestExecute_passwordDoesNotExistsInFrame_Success(t *testing.T) {
	username := "myUsername"
	passVar := "password"
	password := "myPassword"

	credentialsVar := "myCredentials"

	instruction := retrieves.NewRetrieveForTests(passVar, credentialsVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				credentialsVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithCredentialsForTests(
						credentials.NewCredentialsForTests(username, []byte(password)),
					),
				),
			),
		}),
	)

	account := accounts.NewAccountForTests(username, encryptors.NewEncryptorForTests(4096))
	repository := mocks.NewAccountRepositoryWithRetrieveForTests(map[string]accounts.Account{
		username: account,
	})

	application := NewApplication(repository)
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
	if code != failures.CouldNotFetchPasswordFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchPasswordFromFrame, code)
		return
	}
}
