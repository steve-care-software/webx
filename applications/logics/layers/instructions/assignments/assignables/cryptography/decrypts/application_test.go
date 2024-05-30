package decrypts

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
	"github.com/steve-care-software/datastencil/domain/stacks/mocks"
)

func TestExecute_withDecrypt_cipherExistsInFrame_passwordExistsInFrame_decryptSucceeds_Success(t *testing.T) {
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

	instruction := decrypts.NewDecryptForTests(cipherVar, passwordVar)
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{
			string(cipher): map[string][]byte{
				string(password): message,
			},
		},
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

	retMessage := retAssignable.Bytes()
	if !bytes.Equal(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestExecute_withDecrypt_cipherExistsInFrame_passwordExistsInFrame_decryptFails_returnsError(t *testing.T) {
	cipherVar := "cipherVar"
	cipher := []byte("this is a cipher")
	passwordVar := "passVar"
	password := []byte("this is a password")

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

	instruction := decrypts.NewDecryptForTests(cipherVar, passwordVar)
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
	if code != failures.CouldNotDecryptCipher {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotDecryptCipher, code)
		return
	}
}

func TestExecute_withDecrypt_cipherExistsInFrame_passwordDoesNotExistsInFrame_returnsError(t *testing.T) {
	cipherVar := "cipherVar"
	cipher := []byte("this is a cipher")
	passwordVar := "passVar"

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

	instruction := decrypts.NewDecryptForTests(cipherVar, passwordVar)
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

func TestExecute_withDecrypt_cipherDoesNotExistsInFrame_returnsError(t *testing.T) {
	cipherVar := "cipherVar"
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

	instruction := decrypts.NewDecryptForTests(cipherVar, passwordVar)
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
	if code != failures.CouldNotFetchCipherFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchCipherFromFrame, code)
		return
	}
}
