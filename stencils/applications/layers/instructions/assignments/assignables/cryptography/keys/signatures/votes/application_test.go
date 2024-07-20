package votes

import (
	"testing"

	"github.com/steve-care-software/datastencil/states/domain/hash"
	application_creates "github.com/steve-care-software/datastencil/stencils/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	application_validates "github.com/steve-care-software/datastencil/stencils/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/datastencil/stencils/domain/keys/signers"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
)

func TestExecute_withCreate_Success(t *testing.T) {
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

	instruction := votes.NewVoteWithCreateForTests(
		creates.NewCreateForTests(messageVar, ringVar, pkVar),
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

func TestExecute_withValidate_Success(t *testing.T) {
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

	instruction := votes.NewVoteWithValidateForTests(
		validates.NewValidateForTests(voteVar, messageVar, ringVar),
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
