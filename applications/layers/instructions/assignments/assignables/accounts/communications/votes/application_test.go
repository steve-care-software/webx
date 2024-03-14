package votes

import (
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_messageExistsInFrame_ringExistsInFrame_accountExistsInFrame_Success(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a message")
	ringVar := "myRingVar"
	ring := signers.NewRingForTests(20)
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	vote := votes.NewVoteForTests(messageVar, ringVar, accountVar)
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

	application := NewApplication()
	retAssignable, pCode, err := application.Execute(frame, vote)
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

func TestExecute_messageDoesNotExistsInFrame_ReturnsError(t *testing.T) {
	messageVar := "myMessage"
	ringVar := "myRingVar"
	ring := signers.NewRingForTests(20)
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	vote := votes.NewVoteForTests(messageVar, ringVar, accountVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
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

	application := NewApplication()
	_, pCode, err := application.Execute(frame, vote)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchMessageFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchMessageFromFrame, code)
		return
	}
}

func TestExecute_messageExistsInFrame_ringDoesNotExistsInFrame_returnsError(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a message")
	ringVar := "myRingVar"
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	vote := votes.NewVoteForTests(messageVar, ringVar, accountVar)
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

	application := NewApplication()
	_, pCode, err := application.Execute(frame, vote)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchRingFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchRingFromFrame, code)
		return
	}
}

func TestExecute_messageExistsInFrame_ringExistsInFrame_accountDoesNotExistsInFrame_ReturnsError(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a message")
	ringVar := "myRingVar"
	ring := signers.NewRingForTests(20)
	accountVar := "myAccount"

	vote := votes.NewVoteForTests(messageVar, ringVar, accountVar)
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
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, vote)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchAccountFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchAccountFromFrame, code)
		return
	}
}
