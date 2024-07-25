package creates

import (
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/signers"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithSignerPublicKeyForTests(
			firstPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			secondPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			thirdPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			signer.PublicKey(),
		),
	})

	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromString(firstPubKey.String())
	secondHash, _ := hashAdapter.FromString(secondPubKey.String())
	thirdHash, _ := hashAdapter.FromString(thirdPubKey.String())
	signerPubKeyHash, _ := hashAdapter.FromString(signer.PublicKey().String())
	pubKeyHashes := []hash.Hash{
		*firstHash,
		*secondHash,
		*thirdHash,
		*signerPubKeyHash,
	}

	messageVar := "myMessage"
	ringVar := "myVar"
	pkVar := "myPK"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithListForTests(ring),
			),
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithSignerForTests(signer),
			),
		}),
	)

	instruction := creates.NewCreateForTests(messageVar, ringVar, pkVar)
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

	if !retAssignable.IsVote() {
		t.Errorf("the assignable was expected to contain vote")
		return
	}

	retVote := retAssignable.Vote()
	if !retVote.Verify(string(message)) {
		t.Errorf("the returned vote could not be verified")
		return
	}

	isVerified, err := signers.NewVoteAdapter().ToVerification(retVote, string(message), pubKeyHashes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !isVerified {
		t.Errorf("the vote was expected to be verified")
		return
	}
}

func TestExecute_messageNotInFrame_returnsError(t *testing.T) {
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithSignerPublicKeyForTests(
			firstPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			secondPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			thirdPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			signer.PublicKey(),
		),
	})

	messageVar := "myMessage"
	ringVar := "myVar"
	pkVar := "myPK"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithListForTests(ring),
			),
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithSignerForTests(signer),
			),
		}),
	)

	instruction := creates.NewCreateForTests(messageVar, ringVar, pkVar)
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

func TestExecute_ringNotInFrame_returnsError(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()

	messageVar := "myMessage"
	ringVar := "myVar"
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

	instruction := creates.NewCreateForTests(messageVar, ringVar, pkVar)
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

	if *pCode != failures.CouldNotFetchRingFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchRingFromFrame)
		return
	}
}

func TestExecute_signerNotInFrame_returnsError(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithSignerPublicKeyForTests(
			firstPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			secondPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			thirdPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			signer.PublicKey(),
		),
	})

	messageVar := "myMessage"
	ringVar := "myVar"
	pkVar := "myPK"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithListForTests(ring),
			),
		}),
	)

	instruction := creates.NewCreateForTests(messageVar, ringVar, pkVar)
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

func TestExecute_signerPublicKeyNotInRing_returnsError(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithSignerPublicKeyForTests(
			firstPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			secondPubKey,
		),
		stacks.NewAssignableWithSignerPublicKeyForTests(
			thirdPubKey,
		),
	})

	messageVar := "myMessage"
	ringVar := "myVar"
	pkVar := "myPK"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithListForTests(ring),
			),
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithSignerForTests(signer),
			),
		}),
	)

	instruction := creates.NewCreateForTests(messageVar, ringVar, pkVar)
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

	if *pCode != failures.CouldNotVoteOnMessageInFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotVoteOnMessageInFrame)
		return
	}
}
