package assignments

import (
	"bytes"
	"testing"

	application_assignables "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables"
	application_bytes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/bytes"
	application_compilers "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/compilers"
	application_constants "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/constants"
	application_cryptography "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography"
	application_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/decrypts"
	application_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables"
	assignable_bytes "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_Success(t *testing.T) {
	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				firstVariable,
				stacks.NewAssignableWithBytesForTests(
					firstValue,
				),
			),
			stacks.NewAssignmentForTests(
				secondVariable,
				stacks.NewAssignableWithBytesForTests(
					secondValue,
				),
			),
		}),
	)

	name := "myName"
	instruction := assignments.NewAssignmentForTests(
		name,
		assignables.NewAssignableWithBytesForTests(
			assignable_bytes.NewBytesWithJoinForTests([]string{
				firstVariable,
				secondVariable,
			}),
		),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_assignables.NewApplication(
			application_compilers.NewApplication(
				mocks.NewInstanceAdapter(
					map[string][]byte{},
					map[string]instances.Instance{},
				),
			),
			application_bytes.NewApplication(),
			application_constants.NewApplication(),
			application_cryptography.NewApplication(
				application_decrypts.NewApplication(
					encryptor,
				),
				application_encrypts.NewApplication(
					encryptor,
				),
				nil,
			),
		),
	)

	retAssignment, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	retName := retAssignment.Name()
	if name != retName {
		t.Errorf("the name was expected to be '%s', '%s' returned", name, retName)
		return
	}

	retAssignable := retAssignment.Assignable()
	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retBytes := retAssignable.Bytes()
	expected := bytes.Join([][]byte{
		firstValue,
		secondValue,
	}, []byte{})

	if !bytes.Equal(expected, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}
