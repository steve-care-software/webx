package encryptions

import (
	"bytes"
	"reflect"
	"testing"

	application_decrypts "github.com/steve-care-software/webx/engine/stencils/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	application_encrypts "github.com/steve-care-software/webx/engine/stencils/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	"github.com/steve-care-software/webx/engine/stencils/domain/keys/encryptors"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks/failures"
)

func TestExecute_withEncrypt_Success(t *testing.T) {
	message := []byte("this is some data")
	encryptor := encryptors.NewEncryptorForTests(1048)
	pubKey := encryptor.Public()

	messageVar := "myMessage"
	pubKeyVar := "myCipher"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				pubKeyVar,
				stacks.NewAssignableWithEncryptorPublicKeyForTests(pubKey),
			),
		}),
	)

	instruction := encryptions.NewEncryptionWithEncryptForTests(
		encrypts.NewEncryptForTests(messageVar, pubKeyVar),
	)

	application := NewApplication(
		application_decrypts.NewApplication(),
		application_encrypts.NewApplication(),
		1048,
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
	retMessage, err := encryptor.Decrypt(retCipher)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestExecute_withDecrypt_Success(t *testing.T) {
	data := []byte("this is some data")
	encryptor := encryptors.NewEncryptorForTests(1048)
	cipher, err := encryptor.Public().Encrypt(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cipherVar := "myCipher"
	pkVar := "myPK"

	instruction := encryptions.NewEncryptionWithDecryptForTests(
		decrypts.NewDecryptForTests(cipherVar, pkVar),
	)

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

	application := NewApplication(
		application_decrypts.NewApplication(),
		application_encrypts.NewApplication(),
		1048,
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

	retBytes := retAssignable.Bytes()
	if !bytes.Equal(data, retBytes) {
		t.Errorf("the returned bytes are invalid")
		return
	}
}

func TestExecute_withFetchPublicKey_Success(t *testing.T) {
	encryptor := encryptors.NewEncryptorForTests(1048)

	pkVar := "myPK"
	instruction := encryptions.NewEncryptionWithFetchPublicKeyForTests(
		pkVar,
	)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithEncryptorForTests(encryptor),
			),
		}),
	)

	application := NewApplication(
		application_decrypts.NewApplication(),
		application_encrypts.NewApplication(),
		1048,
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

	if !retAssignable.IsEncryptorPublicKey() {
		t.Errorf("the assignable was expected to contain an encryptor public key")
		return
	}

	retEncryptorPubKey := retAssignable.EncryptorPublicKey()
	if !reflect.DeepEqual(encryptor.Public(), retEncryptorPubKey) {
		t.Errorf("the returned public key is invalid")
		return
	}
}

func TestExecute_fetchPublicKey_withoutPKInFrame_returnsError(t *testing.T) {
	pkVar := "myPK"
	instruction := encryptions.NewEncryptionWithFetchPublicKeyForTests(
		pkVar,
	)

	frame := stacks.NewFrameForTests()

	application := NewApplication(
		application_decrypts.NewApplication(),
		application_encrypts.NewApplication(),
		1048,
	)

	_, pCode, err := application.Execute(frame, instruction)
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

func TestExecute_withGeneratePK_Success(t *testing.T) {
	instruction := encryptions.NewEncryptionWithGeneratePrivateKeyForTests()
	frame := stacks.NewFrameForTests()

	application := NewApplication(
		application_decrypts.NewApplication(),
		application_encrypts.NewApplication(),
		1048,
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

	if !retAssignable.IsEncryptor() {
		t.Errorf("the assignable was expected to contain an encryptor")
		return
	}
}
