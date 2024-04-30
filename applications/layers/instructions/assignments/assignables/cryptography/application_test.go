package cryptography

import (
	"bytes"
	"testing"

	application_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/decrypts"
	application_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_withDecrypt_Success(t *testing.T) {
	cipherVar := "cipherVar"
	cipher := []byte("this is a cipher")
	passwordVar := "passVar"
	password := []byte("this is a password")
	message := []byte("this is some message")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests(
					cipher,
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

	instruction := cryptography.NewCryptographyWithDecryptForTests(
		decrypts.NewDecryptForTests(cipherVar, passwordVar),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{
			string(cipher): map[string][]byte{
				string(password): message,
			},
		},
	)

	application := NewApplication(
		application_decrypts.NewApplication(
			encryptor,
		),
		application_encrypts.NewApplication(
			encryptor,
		),
		nil,
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
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestExecute_withEncrypt_Success(t *testing.T) {
	messageVar := "messageVar"
	message := []byte("this is a message")
	passwordVar := "passVar"
	password := []byte("this is a password")
	cipher := []byte("this is some cipher")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(
					message,
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

	instruction := cryptography.NewCryptographyWithEncryptForTests(
		encrypts.NewEncryptForTests(messageVar, passwordVar),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{
			string(message): map[string][]byte{
				string(password): cipher,
			},
		},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_decrypts.NewApplication(
			encryptor,
		),
		application_encrypts.NewApplication(
			encryptor,
		),
		nil,
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
	if !bytes.Equal(cipher, retCipher) {
		t.Errorf("the returned cipher is invalid")
		return
	}
}
