package creates

import (
	"testing"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/stencils/domain/keys/signers"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
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

	instruction := creates.NewCreateForTests(messageVar, pkVar)
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

func TestExecute_messageNotInFrame_returnsError(t *testing.T) {
	signer := signers.NewFactory().Create()

	messageVar := "myMessage"
	pkVar := "myPK"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithSignerForTests(signer),
			),
		}),
	)

	instruction := creates.NewCreateForTests(messageVar, pkVar)
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

func TestExecute_pkNotInFrame_returnsError(t *testing.T) {
	message := []byte("this is a message")
	messageVar := "myMessage"
	pkVar := "myPK"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
		}),
	)

	instruction := creates.NewCreateForTests(messageVar, pkVar)
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
