package credentials

import (
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_usernameExistsInFrame_passwordExistsInFrame_Success(t *testing.T) {
	usernameVar := "usernameVar"
	username := "myUsername"
	passwordVar := "passwordVar"
	password := []byte("myPassword")

	instruction := credentials.NewCredentialsForTests(usernameVar, passwordVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				usernameVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(username),
				),
			),
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					password,
				),
			),
		}),
	)

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

	if !retAssignable.IsAccount() {
		t.Errorf("the assignable was expected to contain an Account instance")
		return
	}
}

func TestExecute_usernameExistsInFrame_passwordDoesNotExistsInFrame_ReturnsError(t *testing.T) {
	usernameVar := "usernameVar"
	username := "myUsername"
	passwordVar := "passwordVar"
	instruction := credentials.NewCredentialsForTests(usernameVar, passwordVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				usernameVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(username),
				),
			),
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchPasswordFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchPasswordFromFrame, code)
		return
	}
}

func TestExecute_usernameDoesNotExistsInFrame_ReturnsError(t *testing.T) {
	usernameVar := "usernameVar"
	passwordVar := "passwordVar"
	password := []byte("myPassword")

	instruction := credentials.NewCredentialsForTests(usernameVar, passwordVar)
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					password,
				),
			),
		}),
	)

	application := NewApplication()
	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchUsernameFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchUsernameFromFrame, code)
		return
	}
}
