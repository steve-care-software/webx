package signs

import (
	"testing"

	application_creates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	application_validates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/signers"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

func TestExecute_withCreate_Success(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()

	messageVar := "myMessage"
	pkVar := "myPK"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithSignerForTests(signer),
			),
		}),
	)

	instruction := signs.NewSignWithCreateForTests(
		creates.NewCreateForTests(messageVar, pkVar),
	)

	application := NewApplication(
		application_creates.NewApplication(),
		application_validates.NewApplication(),
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

	if !retAssignable.IsSignature() {
		t.Errorf("the assignable was expected to contain signature")
		return
	}

	retSignature := retAssignable.Signature()
	retPubKey := retSignature.PublicKey(messageVar)
	if retPubKey.Equals(signer.PublicKey()) {
		t.Errorf("the returned public key is invalid")
		return
	}
}

func TestExecute_withValidate_Success(t *testing.T) {
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

	instruction := signs.NewSignWithValidateForTests(
		validates.NewValidateForTests(signatureVar, messageVar, pubKeyVar),
	)

	application := NewApplication(
		application_creates.NewApplication(),
		application_validates.NewApplication(),
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
