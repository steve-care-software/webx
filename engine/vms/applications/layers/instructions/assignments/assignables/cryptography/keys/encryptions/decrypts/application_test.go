package decrypts

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/encryptors"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {

	data := []byte("this is some data")
	encryptor := encryptors.NewEncryptorForTests(1048)
	cipher, err := encryptor.Public().Encrypt(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cipherVar := "myCipher"
	pkVar := "myPK"
	decrypt := decrypts.NewDecryptForTests(cipherVar, pkVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests(cipher),
			),
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithEncryptorForTests(encryptor),
			),
		}),
	)

	application := NewApplication()
	retAssignable, pCode, err := application.Execute(frame, decrypt)
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

	retBytes := retAssignable.Bytes()
	if !bytes.Equal(data, retBytes) {
		t.Errorf("the returned bytes are invalid")
		return
	}
}

func TestExecute_cipherAssignmentDoesNotExistsInFrame_returnsError(t *testing.T) {
	encryptor := encryptors.NewEncryptorForTests(1048)
	cipherVar := "myCipher"
	pkVar := "myPK"
	decrypt := decrypts.NewDecryptForTests(cipherVar, pkVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithEncryptorForTests(encryptor),
			),
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, decrypt)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchCipherFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchCipherFromFrame)
		return
	}
}

func TestExecute_pkAssignmentDoesNotExistsInFrame_returnsError(t *testing.T) {
	data := []byte("this is some data")
	encryptor := encryptors.NewEncryptorForTests(1048)
	cipher, err := encryptor.Public().Encrypt(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cipherVar := "myCipher"
	pkVar := "myPK"
	decrypt := decrypts.NewDecryptForTests(cipherVar, pkVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests(cipher),
			),
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, decrypt)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchEncryptionPrivateKeyFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchEncryptionPrivateKeyFromFrame)
		return
	}
}

func TestExecute_invalidCipher_returnsError(t *testing.T) {
	encryptor := encryptors.NewEncryptorForTests(1048)
	cipherVar := "myCipher"
	pkVar := "myPK"
	decrypt := decrypts.NewDecryptForTests(cipherVar, pkVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests([]byte("this is an invalid cipher")),
			),
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithEncryptorForTests(encryptor),
			),
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, decrypt)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotDecryptCipher {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotDecryptCipher)
		return
	}
}
