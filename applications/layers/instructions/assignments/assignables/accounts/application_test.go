package accounts

import (
	"bytes"
	"reflect"
	"testing"

	application_execution_communications "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications"
	application_signs "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/signs"
	application_votes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/votes"
	application_execution_credentials "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/credentials"
	application_execution_encryptions "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions"
	application_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	application_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	application_execution_retrieves "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/accounts"
	account_credentials "github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	assignables_accounts "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stack_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_withList_passwordExistsInFrame_accountsListCanBeRetrievedFromDatabase_Success(t *testing.T) {
	usernameList := []string{
		"first",
		"second",
	}

	passwordVar := "passwordVar"
	password := "myPassword"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(password),
				),
			),
		}),
	)

	instruction := assignables_accounts.NewAccountWithListForTests(passwordVar)

	repository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{
		password: usernameList,
	})

	application := NewApplication(
		application_execution_communications.NewApplication(
			application_signs.NewApplication(),
			application_votes.NewApplication(),
		),
		application_execution_credentials.NewApplication(),
		application_execution_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
		),
		application_execution_retrieves.NewApplication(
			repository,
		),
		repository,
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

	if !retAssignable.IsStringList() {
		t.Errorf("the assignable was expected to contain a string list")
		return
	}

	retStringList := retAssignable.StringList()
	if !reflect.DeepEqual(usernameList, retStringList) {
		t.Errorf("the returned usernames list is invalid")
		return
	}
}

func TestExecute_withList_passwordExistsInFrame_accountsListCanNotBeRetrievedFromDatabase_returnsError(t *testing.T) {
	passwordVar := "passwordVar"
	password := "myPassword"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(password),
				),
			),
		}),
	)

	instruction := assignables_accounts.NewAccountWithListForTests(passwordVar)

	repository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})

	application := NewApplication(
		application_execution_communications.NewApplication(
			application_signs.NewApplication(),
			application_votes.NewApplication(),
		),
		application_execution_credentials.NewApplication(),
		application_execution_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
		),
		application_execution_retrieves.NewApplication(
			repository,
		),
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
	if code != failures.CouldNotRetrieveAccountNamesListFromDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotRetrieveAccountNamesListFromDatabase, code)
		return
	}
}

func TestExecute_withList_passwordDoesNotExistsInFrame_ReturnsError(t *testing.T) {
	usernameList := []string{
		"first",
		"second",
	}

	passwordVar := "passwordVar"
	password := "myPassword"

	frame := stacks.NewFrameForTests()
	instruction := assignables_accounts.NewAccountWithListForTests(passwordVar)

	repository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{
		password: usernameList,
	})

	application := NewApplication(
		application_execution_communications.NewApplication(
			application_signs.NewApplication(),
			application_votes.NewApplication(),
		),
		application_execution_credentials.NewApplication(),
		application_execution_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
		),
		application_execution_retrieves.NewApplication(
			repository,
		),
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
	if code != failures.CouldNotFetchPasswordFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchPasswordFromFrame, code)
		return
	}
}

func TestExecute_withCredentials_Success(t *testing.T) {
	usernameVar := "usernameVar"
	username := "myUsername"
	passwordVar := "passwordVar"
	password := []byte("myPassword")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				usernameVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(username),
				),
			),
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					password,
				),
			),
		}),
	)

	instruction := assignables_accounts.NewAccountWithCredentialsForTests(
		credentials.NewCredentialsForTests(usernameVar, passwordVar),
	)

	repository := mocks.NewAccountRepositoryWithRetrieveForTests(map[string]accounts.Account{})

	application := NewApplication(
		application_execution_communications.NewApplication(
			application_signs.NewApplication(),
			application_votes.NewApplication(),
		),
		application_execution_credentials.NewApplication(),
		application_execution_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
		),
		application_execution_retrieves.NewApplication(
			repository,
		),
		repository,
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

	if !retAssignable.IsAccount() {
		t.Errorf("the assignable was expected to contain an Account")
		return
	}

	retAccount := retAssignable.Account()
	if !retAccount.IsCredentials() {
		t.Errorf("the assignable (account) was expected to contain a Credentials")
		return
	}
}

func TestExecute_withRetrieve_Success(t *testing.T) {
	username := "myUsername"
	passVar := "password"
	password := "myPassword"

	credentialsVar := "myCredentials"

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
						account_credentials.NewCredentialsForTests(username, []byte(password)),
					),
				),
			),
		}),
	)

	instruction := assignables_accounts.NewAccountWithRetrieveForTests(
		retrieves.NewRetrieveForTests(passVar, credentialsVar),
	)

	account := accounts.NewAccountForTests(username, encryptors.NewEncryptorForTests(4096))
	repository := mocks.NewAccountRepositoryWithRetrieveForTests(map[string]accounts.Account{
		username: account,
	})

	application := NewApplication(
		application_execution_communications.NewApplication(
			application_signs.NewApplication(),
			application_votes.NewApplication(),
		),
		application_execution_credentials.NewApplication(),
		application_execution_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
		),
		application_execution_retrieves.NewApplication(
			repository,
		),
		repository,
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

func TestExecute_withCommunication_Success(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a message")
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(
					message,
				),
			),
			stacks.NewAssignmentForTests(
				accountVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithAccountForTests(
						account,
					),
				),
			),
		}),
	)

	instruction := assignables_accounts.NewAccountWithCommunicationForTests(
		communications.NewCommunicationWithSignForTests(
			signs.NewSignForTests(messageVar, accountVar),
		),
	)

	repository := mocks.NewAccountRepositoryWithRetrieveForTests(map[string]accounts.Account{})

	application := NewApplication(
		application_execution_communications.NewApplication(
			application_signs.NewApplication(),
			application_votes.NewApplication(),
		),
		application_execution_credentials.NewApplication(),
		application_execution_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
		),
		application_execution_retrieves.NewApplication(
			repository,
		),
		repository,
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

	if !retAssignable.IsAccount() {
		t.Errorf("the assignable was expected to contain an Account instance")
		return
	}

	retAccount := retAssignable.Account()
	if !retAccount.IsSignature() {
		t.Errorf("the assignable (account) was expected to contain a Signature instance")
		return
	}

	retSignature := retAccount.Signature()
	sigPubKey := retSignature.PublicKey(string(message))
	expectedPubKey := account.Signer().PublicKey()

	if !sigPubKey.Equals(expectedPubKey) {
		t.Errorf("the signature's public key is invalid")
		return
	}
}

func TestExecute_withEncryption_Success(t *testing.T) {
	message := []byte("this is a cipher")
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	cipherVar := "cipherVar"
	cipher, err := account.Encryptor().Public().Encrypt(message)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests(
					cipher,
				),
			),
			stacks.NewAssignmentForTests(
				accountVar,
				stacks.NewAssignableWithAccountForTests(
					stack_accounts.NewAccountWithAccountForTests(
						account,
					),
				),
			),
		}),
	)

	instruction := assignables_accounts.NewAccountWithEncryptionForTests(
		encryptions.NewEncryptionWithDecryptForTests(
			decrypts.NewDecryptForTests(cipherVar, accountVar),
		),
	)

	repository := mocks.NewAccountRepositoryWithRetrieveForTests(map[string]accounts.Account{})

	application := NewApplication(
		application_execution_communications.NewApplication(
			application_signs.NewApplication(),
			application_votes.NewApplication(),
		),
		application_execution_credentials.NewApplication(),
		application_execution_encryptions.NewApplication(
			application_decrypts.NewApplication(),
			application_encrypts.NewApplication(),
		),
		application_execution_retrieves.NewApplication(
			repository,
		),
		repository,
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

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retMessage := retAssignable.Bytes()
	if !bytes.Equal(message, retMessage) {
		t.Errorf("the cipher could not be decrypted")
		return
	}
}
