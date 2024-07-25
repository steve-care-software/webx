package validates

import (
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/signers"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	signature, err := signer.Sign(string(message))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	publicKey := signer.PublicKey()

	signatureVar := "mySignature"
	messageVar := "myMessage"
	pubKeyVar := "myPubKey"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				signatureVar,
				stacks.NewAssignableWithSignatureForTests(signature),
			),
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				pubKeyVar,
				stacks.NewAssignableWithSignerPublicKeyForTests(publicKey),
			),
		}),
	)

	instruction := validates.NewValidateForTests(signatureVar, messageVar, pubKeyVar)
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

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain a bool")
		return
	}

	pBool := retAssignable.Bool()
	if !*pBool {
		t.Errorf("the returned boolean was expected to be true")
		return
	}
}

func TestExecute_couldNotVerify_Success(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	signature, err := signers.NewFactory().Create().Sign(string(message))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	publicKey := signer.PublicKey()

	signatureVar := "mySignature"
	messageVar := "myMessage"
	pubKeyVar := "myPubKey"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				signatureVar,
				stacks.NewAssignableWithSignatureForTests(signature),
			),
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				pubKeyVar,
				stacks.NewAssignableWithSignerPublicKeyForTests(publicKey),
			),
		}),
	)

	instruction := validates.NewValidateForTests(signatureVar, messageVar, pubKeyVar)
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

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain a bool")
		return
	}

	pBool := retAssignable.Bool()
	if *pBool {
		t.Errorf("the returned boolean was expected to be false")
		return
	}
}

func TestExecute_signatureNotInFrame_returnsError(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()

	publicKey := signer.PublicKey()

	signatureVar := "mySignature"
	messageVar := "myMessage"
	pubKeyVar := "myPubKey"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				pubKeyVar,
				stacks.NewAssignableWithSignerPublicKeyForTests(publicKey),
			),
		}),
	)

	instruction := validates.NewValidateForTests(signatureVar, messageVar, pubKeyVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchSignatureFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchSignatureFromFrame)
		return
	}
}

func TestExecute_messageNotInFrame_returnsError(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	signature, err := signer.Sign(string(message))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	publicKey := signer.PublicKey()

	signatureVar := "mySignature"
	messageVar := "myMessage"
	pubKeyVar := "myPubKey"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				signatureVar,
				stacks.NewAssignableWithSignatureForTests(signature),
			),
			stacks.NewAssignmentForTests(
				pubKeyVar,
				stacks.NewAssignableWithSignerPublicKeyForTests(publicKey),
			),
		}),
	)

	instruction := validates.NewValidateForTests(signatureVar, messageVar, pubKeyVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
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

func TestExecute_pubKeyNotInFrame_returnsError(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	signature, err := signer.Sign(string(message))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	signatureVar := "mySignature"
	messageVar := "myMessage"
	pubKeyVar := "myPubKey"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				signatureVar,
				stacks.NewAssignableWithSignatureForTests(signature),
			),
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
		}),
	)

	instruction := validates.NewValidateForTests(signatureVar, messageVar, pubKeyVar)
	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if *pCode != failures.CouldNotFetchSignerPrivateKeyFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchSignerPrivateKeyFromFrame)
		return
	}
}
