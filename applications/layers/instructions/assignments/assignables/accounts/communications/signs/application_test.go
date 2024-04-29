package signs

import (
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

func TestExecute_messageExistsInFrame_accountExistsInFrame_Success(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a message")
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	sign := signs.NewSignForTests(messageVar, accountVar)
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
	retAssignable, pCode, err := application.Execute(frame, sign)
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

func TestExecute_messageDoesNotExistsInFrame_returnsError(t *testing.T) {
	messageVar := "myMessage"
	accountVar := "myAccount"
	account := accounts.NewAccountForTests("myUsername", encryptors.NewEncryptorForTests(4096))

	sign := signs.NewSignForTests(messageVar, accountVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
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
	_, pCode, err := application.Execute(frame, sign)
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

func TestExecute_messageExistsInFrame_accountDoesNotExistsInFrame_ReturnsError(t *testing.T) {
	messageVar := "myMessage"
	message := []byte("this is a message")
	accountVar := "myAccount"
	sign := signs.NewSignForTests(messageVar, accountVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				messageVar,
				stacks.NewAssignableWithBytesForTests(
					message,
				),
			),
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, sign)
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
