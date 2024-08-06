package validates

import (
	"testing"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/signers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

func TestExecute_Success(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	pubKeys := []signers.PublicKey{
		firstPubKey,
		secondPubKey,
		thirdPubKey,
		signer.PublicKey(),
	}

	vote, err := signer.Vote(string(message), pubKeys)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromString(firstPubKey.String())
	secondHash, _ := hashAdapter.FromString(secondPubKey.String())
	thirdHash, _ := hashAdapter.FromString(thirdPubKey.String())
	signerPubKeyHash, _ := hashAdapter.FromString(signer.PublicKey().String())
	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithHashForTests(
			*firstHash,
		),
		stacks.NewAssignableWithHashForTests(
			*secondHash,
		),
		stacks.NewAssignableWithHashForTests(
			*thirdHash,
		),
		stacks.NewAssignableWithHashForTests(
			*signerPubKeyHash,
		),
	})

	voteVar := "myVote"
	messageVar := "myMessage"
	ringVar := "myRing"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				voteVar,
				stacks.NewAssignableWithVoteForTests(vote),
			),
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

	instruction := validates.NewValidateForTests(voteVar, messageVar, ringVar)
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

func TestExecute_withInvalidVote_Success(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	pubKeys := []signers.PublicKey{
		firstPubKey,
		secondPubKey,
		thirdPubKey,
		signer.PublicKey(),
	}

	vote, err := signer.Vote(string(message), pubKeys)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromString(firstPubKey.String())
	secondHash, _ := hashAdapter.FromString(secondPubKey.String())
	thirdHash, _ := hashAdapter.FromString(thirdPubKey.String())
	signerPubKeyHash, _ := hashAdapter.FromString(signer.PublicKey().String())
	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithHashForTests(
			*firstHash,
		),
		stacks.NewAssignableWithHashForTests(
			*secondHash,
		),
		stacks.NewAssignableWithHashForTests(
			*thirdHash,
		),
		stacks.NewAssignableWithHashForTests(
			*signerPubKeyHash,
		),
	})

	voteVar := "myVote"
	messageVar := "myMessage"
	ringVar := "myRing"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				voteVar,
				stacks.NewAssignableWithVoteForTests(vote),
			),
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests([]byte("this is an invalid message")),
			),
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithListForTests(ring),
			),
		}),
	)

	instruction := validates.NewValidateForTests(voteVar, messageVar, ringVar)
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

func TestExecute_voteIsNotInFrame_returnsError(t *testing.T) {
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromString(firstPubKey.String())
	secondHash, _ := hashAdapter.FromString(secondPubKey.String())
	thirdHash, _ := hashAdapter.FromString(thirdPubKey.String())
	signerPubKeyHash, _ := hashAdapter.FromString(signer.PublicKey().String())
	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithHashForTests(
			*firstHash,
		),
		stacks.NewAssignableWithHashForTests(
			*secondHash,
		),
		stacks.NewAssignableWithHashForTests(
			*thirdHash,
		),
		stacks.NewAssignableWithHashForTests(
			*signerPubKeyHash,
		),
	})

	voteVar := "myVote"
	messageVar := "myMessage"
	ringVar := "myRing"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests([]byte("this is an invalid message")),
			),
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithListForTests(ring),
			),
		}),
	)

	instruction := validates.NewValidateForTests(voteVar, messageVar, ringVar)
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

	if *pCode != failures.CouldNotFetchVoteFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchVoteFromFrame)
		return
	}
}

func TestExecute_messageIsNotInFrame_returnsError(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	pubKeys := []signers.PublicKey{
		firstPubKey,
		secondPubKey,
		thirdPubKey,
		signer.PublicKey(),
	}

	vote, err := signer.Vote(string(message), pubKeys)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromString(firstPubKey.String())
	secondHash, _ := hashAdapter.FromString(secondPubKey.String())
	thirdHash, _ := hashAdapter.FromString(thirdPubKey.String())
	signerPubKeyHash, _ := hashAdapter.FromString(signer.PublicKey().String())
	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithHashForTests(
			*firstHash,
		),
		stacks.NewAssignableWithHashForTests(
			*secondHash,
		),
		stacks.NewAssignableWithHashForTests(
			*thirdHash,
		),
		stacks.NewAssignableWithHashForTests(
			*signerPubKeyHash,
		),
	})

	voteVar := "myVote"
	messageVar := "myMessage"
	ringVar := "myRing"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				voteVar,
				stacks.NewAssignableWithVoteForTests(vote),
			),
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithListForTests(ring),
			),
		}),
	)

	instruction := validates.NewValidateForTests(voteVar, messageVar, ringVar)
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
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	pubKeys := []signers.PublicKey{
		firstPubKey,
		secondPubKey,
		thirdPubKey,
		signer.PublicKey(),
	}

	vote, err := signer.Vote(string(message), pubKeys)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	voteVar := "myVote"
	messageVar := "myMessage"
	ringVar := "myRing"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				voteVar,
				stacks.NewAssignableWithVoteForTests(vote),
			),
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(message),
			),
		}),
	)

	instruction := validates.NewValidateForTests(voteVar, messageVar, ringVar)
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

func TestExecute_elementInRingNotAnHash_returnsError(t *testing.T) {
	message := []byte("this is a message")
	signer := signers.NewFactory().Create()
	firstPubKey := signers.NewFactory().Create().PublicKey()
	secondPubKey := signers.NewFactory().Create().PublicKey()
	thirdPubKey := signers.NewFactory().Create().PublicKey()

	pubKeys := []signers.PublicKey{
		firstPubKey,
		secondPubKey,
		thirdPubKey,
		signer.PublicKey(),
	}

	vote, err := signer.Vote(string(message), pubKeys)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromString(firstPubKey.String())
	secondHash, _ := hashAdapter.FromString(secondPubKey.String())
	thirdHash, _ := hashAdapter.FromString(thirdPubKey.String())
	signerPubKeyHash, _ := hashAdapter.FromString(signer.PublicKey().String())
	ring := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithHashForTests(
			*firstHash,
		),
		stacks.NewAssignableWithHashForTests(
			*secondHash,
		),
		stacks.NewAssignableWithHashForTests(
			*thirdHash,
		),
		stacks.NewAssignableWithHashForTests(
			*signerPubKeyHash,
		),
		stacks.NewAssignableWithBytesForTests(
			[]byte("this is some data"),
		),
	})

	voteVar := "myVote"
	messageVar := "myMessage"
	ringVar := "myRing"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				voteVar,
				stacks.NewAssignableWithVoteForTests(vote),
			),
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

	instruction := validates.NewValidateForTests(voteVar, messageVar, ringVar)
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

	if *pCode != failures.CouldNotFetchHashFromList {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchHashFromList)
		return
	}
}
