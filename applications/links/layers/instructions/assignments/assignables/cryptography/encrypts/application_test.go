package encrypts

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_withEncrypt_messageExistsInFrame_passwordExistsInFrame_encryptSucceeds_Success(t *testing.T) {
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

	instruction := encrypts.NewEncryptForTests(messageVar, passwordVar)
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{
			string(message): map[string][]byte{
				string(password): cipher,
			},
		},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		encryptor,
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

func TestExecute_withEncrypt_messageExistsInFrame_passwordExistsInFrame_encryptFails_returnsError(t *testing.T) {
	messageVar := "messageVar"
	message := []byte("this is a message")
	passwordVar := "passVar"
	password := []byte("this is a password")

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

	instruction := encrypts.NewEncryptForTests(messageVar, passwordVar)
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		encryptor,
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
	if code != failures.CouldNotEncryptMessage {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotEncryptMessage, code)
		return
	}
}

func TestExecute_withEncrypt_messageExistsInFrame_passwordDoesNotExistsInFrame_returnsError(t *testing.T) {
	messageVar := "messageVar"
	message := []byte("this is a message")
	passwordVar := "passVar"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(
					message,
				),
			),
		}),
	)

	instruction := encrypts.NewEncryptForTests(messageVar, passwordVar)
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		encryptor,
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

func TestExecute_withEncrypt_messageDoesNotExistsInFrame_returnsError(t *testing.T) {
	messageVar := "messageVar"
	passwordVar := "passVar"
	password := []byte("this is a password")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					password,
				),
			),
		}),
	)

	instruction := encrypts.NewEncryptForTests(messageVar, passwordVar)
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		encryptor,
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
	if code != failures.CouldNotFetchMessageFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchMessageFromFrame, code)
		return
	}
}
