package encrypts

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/encryptors"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
	message := []byte("this is some data")
	encryptor := encryptors.NewEncryptorForTests(1048)
	pubKey := encryptor.Public()

	messageVar := "myMessage"
	pubKeyVar := "myCipher"

	encrypt := encrypts.NewEncryptForTests(messageVar, pubKeyVar)
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

	application := NewApplication()
	retAssignable, pCode, err := application.Execute(frame, encrypt)
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

func TestExecute_messageVarNotInStack_returnsError(t *testing.T) {
	encryptor := encryptors.NewEncryptorForTests(1048)
	pubKey := encryptor.Public()

	messageVar := "myMessage"
	pubKeyVar := "myCipher"

	encrypt := encrypts.NewEncryptForTests(messageVar, pubKeyVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				pubKeyVar,
				stacks.NewAssignableWithEncryptorPublicKeyForTests(pubKey),
			),
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, encrypt)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchMessageFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchMessageFromFrame)
		return
	}
}

func TestExecute_pubKeyVarNotInStack_returnsError(t *testing.T) {
	message := []byte("this is some data")

	messageVar := "myMessage"
	pubKeyVar := "myCipher"

	encrypt := encrypts.NewEncryptForTests(messageVar, pubKeyVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, encrypt)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchEncryptionPublicKeyFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchEncryptionPublicKeyFromFrame)
		return
	}
}
