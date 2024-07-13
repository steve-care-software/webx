package signatures

import (
	"testing"

	application_signs "github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	application_signs_creates "github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	application_signs_validates "github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	application_votes "github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	application_votes_creates "github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	application_votes_validates "github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	signs_creates "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	votes_creates "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
	"github.com/steve-care-software/historydb/domain/hash"
)

func TestExecute_withGeneratePrivateKey_Success(t *testing.T) {
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
			stacks.NewAssignmentForTests(
				pkVar,
				stacks.NewAssignableWithSignerForTests(signer),
			),
		}),
	)

	instruction := signatures.NewSignatureWithGeneratePrivateKeyForTests()

	application := NewApplication(
		application_votes.NewApplication(
			application_votes_creates.NewApplication(),
			application_votes_validates.NewApplication(),
		),
		application_signs.NewApplication(
			application_signs_creates.NewApplication(),
			application_signs_validates.NewApplication(),
		),
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

	if !retAssignable.IsSigner() {
		t.Errorf("the assignable was expected to contain a signer")
		return
	}
}

func TestExecute_withSign_Success(t *testing.T) {
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

	instruction := signatures.NewSignatureWithSignForTests(
		signs.NewSignWithCreateForTests(
			signs_creates.NewCreateForTests(messageVar, pkVar),
		),
	)

	application := NewApplication(
		application_votes.NewApplication(
			application_votes_creates.NewApplication(),
			application_votes_validates.NewApplication(),
		),
		application_signs.NewApplication(
			application_signs_creates.NewApplication(),
			application_signs_validates.NewApplication(),
		),
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

func TestExecute_withVote_Success(t *testing.T) {
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

	instruction := signatures.NewSignatureWithVoteForTests(
		votes.NewVoteWithCreateForTests(
			votes_creates.NewCreateForTests(messageVar, ringVar, pkVar),
		),
	)

	application := NewApplication(
		application_votes.NewApplication(
			application_votes_creates.NewApplication(),
			application_votes_validates.NewApplication(),
		),
		application_signs.NewApplication(
			application_signs_creates.NewApplication(),
			application_signs_validates.NewApplication(),
		),
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

func TestExecute_withFetchPublicKey_Success(t *testing.T) {
	signer := signers.NewFactory().Create()
	signerVar := "mySigner"
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				signerVar,
				stacks.NewAssignableWithSignerForTests(signer),
			),
		}),
	)

	instruction := signatures.NewSignatureWithFetchPublicKeyForTests(signerVar)

	application := NewApplication(
		application_votes.NewApplication(
			application_votes_creates.NewApplication(),
			application_votes_validates.NewApplication(),
		),
		application_signs.NewApplication(
			application_signs_creates.NewApplication(),
			application_signs_validates.NewApplication(),
		),
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

	if !retAssignable.IsSignerPublicKey() {
		t.Errorf("the assignable was expected to contain signature")
		return
	}

	retPubKey := retAssignable.SignerPublicKey()
	if !retPubKey.Equals(signer.PublicKey()) {
		t.Errorf("the returned public key is invalid")
		return
	}
}

func TestExecute_withFetchPublicKey_signerNotInFrame_returnsError(t *testing.T) {
	signerVar := "mySigner"
	frame := stacks.NewFrameForTests()
	instruction := signatures.NewSignatureWithFetchPublicKeyForTests(signerVar)

	application := NewApplication(
		application_votes.NewApplication(
			application_votes_creates.NewApplication(),
			application_votes_validates.NewApplication(),
		),
		application_signs.NewApplication(
			application_signs_creates.NewApplication(),
			application_signs_validates.NewApplication(),
		),
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

	if *pCode != failures.CouldNotFetchSignerPrivateKeyFromFrame {
		t.Errorf("the returned code was expected to be %d, %d returned", *pCode, failures.CouldNotFetchSignerPrivateKeyFromFrame)
		return
	}
}
