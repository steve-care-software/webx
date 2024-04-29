package decrypts

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_cipherExistsInFrame_accountExistsInFrame_decryptSucceeds_Success(t *testing.T) {
	message := []byte("this is a cipher")
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	cipherVar := "cipherVar"
	cipher, err := account.Encryptor().Public().Encrypt(message)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instruction := decrypts.NewDecryptForTests(cipherVar, accountVar)
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
					stacks_accounts.NewAccountWithAccountForTests(
						account,
					),
				),
			),
		}),
	)

	application := NewApplication()
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

func TestExecute_cipherExistsInFrame_accountExistsInFrame_withInvalidCipher_returnsError(t *testing.T) {
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	cipherVar := "cipherVar"
	cipher := []byte("this is an invalid cipher")
	instruction := decrypts.NewDecryptForTests(cipherVar, accountVar)
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
					stacks_accounts.NewAccountWithAccountForTests(
						account,
					),
				),
			),
		}),
	)

	application := NewApplication()
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
	if code != failures.CouldNotDecryptCipher {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotDecryptCipher, code)
		return
	}
}

func TestExecute_cipherExistsInFrame_accountDoesNotExistsInFrame_Success(t *testing.T) {
	message := []byte("this is a cipher")
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	cipherVar := "cipherVar"
	cipher, err := account.Encryptor().Public().Encrypt(message)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instruction := decrypts.NewDecryptForTests(cipherVar, accountVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests(
					cipher,
				),
			),
		}),
	)

	application := NewApplication()
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
	if code != failures.CouldNotFetchAccountFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchAccountFromFrame, code)
		return
	}
}

func TestExecute_cipherDoesNotExistsInFrame_Success(t *testing.T) {
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	cipherVar := "cipherVar"
	instruction := decrypts.NewDecryptForTests(cipherVar, accountVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				accountVar,
				stacks.NewAssignableWithAccountForTests(
					stacks_accounts.NewAccountWithAccountForTests(
						account,
					),
				),
			),
		}),
	)

	application := NewApplication()
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
	if code != failures.CouldNotFetchCipherFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchCipherFromFrame, code)
		return
	}
}
