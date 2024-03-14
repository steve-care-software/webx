package communications

import (
	"testing"

	application_signs "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/signs"
	application_votes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"

	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_withSign_Success(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a message")
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))
	communication := communications.NewCommunicationWithSignForTests(
		signs.NewSignForTests(messageVar, accountVar),
	)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(
					message,
				),
			),
			stacks.NewAssignmentForTests(
				accountVar,
				stacks.NewAssignableWithAccountForTests(
					stacks_accounts.NewAccountWithAccountForTests(
						account,
					),
				),
			),
		}),
	)

	application := NewApplication(
		application_signs.NewApplication(),
		application_votes.NewApplication(),
	)

	retAssignable, pCode, err := application.Execute(frame, communication)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsAccount() {
		t.Errorf("the assignable was expected to contain an Account instance")
		return
	}

	retAccount := retAssignable.Account()
	if !retAccount.IsSignature() {
		t.Errorf("the assignable (account) was expected to contain a Signature instance")
		return
	}

	retSignature := retAccount.Signature()
	sigPubKey := retSignature.PublicKey(string(message))
	expectedPubKey := account.Signer().PublicKey()

	if !sigPubKey.Equals(expectedPubKey) {
		t.Errorf("the signature's public key is invalid")
		return
	}
}

func TestExecute_withVote_Success(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a message")
	ringVar := "myRingVar"
	ring := signers.NewRingForTests(20)
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	communication := communications.NewCommunicationWithVoteForTests(
		votes.NewVoteForTests(messageVar, ringVar, accountVar),
	)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(
					message,
				),
			),
			stacks.NewAssignmentForTests(
				ringVar,
				stacks.NewAssignableWithAccountForTests(
					stacks_accounts.NewAccountWithRingForTests(
						ring,
					),
				),
			),
			stacks.NewAssignmentForTests(
				accountVar,
				stacks.NewAssignableWithAccountForTests(
					stacks_accounts.NewAccountWithAccountForTests(
						account,
					),
				),
			),
		}),
	)

	application := NewApplication(
		application_signs.NewApplication(),
		application_votes.NewApplication(),
	)

	retAssignable, pCode, err := application.Execute(frame, communication)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsAccount() {
		t.Errorf("the assignable was expected to contain an Account instance")
		return
	}

	retAccount := retAssignable.Account()
	if !retAccount.IsVote() {
		t.Errorf("the assignable (account) was expected to contain a Signature instance")
		return
	}

	if !retAccount.Vote().Verify(string(message)) {
		t.Errorf("the returned vote could not be verified")
		return
	}
}

func TestExecute_withGenerateRing_generateRingExistsInFrame_Success(t *testing.T) {
	generateRingVar := "amountKeysInRing"
	generateRing := uint(19)

	communication := communications.NewCommunicationWithGenerateRingForTests(
		generateRingVar,
	)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				generateRingVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					generateRing,
				),
			),
		}),
	)

	application := NewApplication(
		application_signs.NewApplication(),
		application_votes.NewApplication(),
	)

	retAssignable, pCode, err := application.Execute(frame, communication)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsAccount() {
		t.Errorf("the assignable was expected to contain an Account instance")
		return
	}

	retAccount := retAssignable.Account()
	if !retAccount.IsRing() {
		t.Errorf("the assignable (account) was expected to contain a Ring")
		return
	}

	retRing := retAccount.Ring()
	if generateRing != uint(len(retRing)) {
		t.Errorf("the ring was expected to contain %d PublicKey, %d returned", generateRing, len(retRing))
		return
	}
}

func TestExecute_withGenerateRing_generateRingDoesNotExistsInFrame_ReturnsError(t *testing.T) {
	generateRingVar := "amountKeysInRing"
	communication := communications.NewCommunicationWithGenerateRingForTests(
		generateRingVar,
	)

	frame := stacks.NewFrameForTests()
	application := NewApplication(
		application_signs.NewApplication(),
		application_votes.NewApplication(),
	)

	_, pCode, err := application.Execute(frame, communication)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchGenerateRingFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchGenerateRingFromFrame, code)
		return
	}
}
