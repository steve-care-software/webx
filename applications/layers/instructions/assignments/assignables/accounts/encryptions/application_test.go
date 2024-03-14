package encryptions

import (
	"bytes"
	"testing"

	application_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	application_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_withDecrypt_Success(t *testing.T) {
	message := []byte("this is a cipher")
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	cipherVar := "cipherVar"
	cipher, err := account.Encryptor().Public().Encrypt(message)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instruction := encryptions.NewEncryptionWithDecryptForTests(
		decrypts.NewDecryptForTests(cipherVar, accountVar),
	)

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

	application := NewApplication(
		application_decrypts.NewApplication(),
		application_encrypts.NewApplication(),
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

func TestExecute_withEncrypt_Success(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a cipher")
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	instruction := encryptions.NewEncryptionWithEncryptForTests(
		encrypts.NewEncryptForTests(messageVar, accountVar),
	)

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
					stacks_accounts.NewAccountWithAccountForTests(
						account,
					),
				),
			),
		}),
	)

	application := NewApplication(
		application_decrypts.NewApplication(),
		application_encrypts.NewApplication(),
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

	retCipher := retAssignable.Bytes()
	retMessage, err := account.Encryptor().Decrypt(retCipher)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(message, retMessage) {
		t.Errorf("the cipher could not be encrypted")
		return
	}
}
