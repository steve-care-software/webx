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
	application_cryptography_keys "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys"
	application_encryptions "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	application_keys_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	application_keys_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	application_signatures "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	application_signs "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	application_signs_creates "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	application_signs_validates "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	application_votes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	application_votes_creates "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	application_votes_validates "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	application_lists "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/lists"
	application_fetches "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables"
	assignable_bytes "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/mocks"
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
				application_cryptography_keys.NewApplication(
					application_encryptions.NewApplication(
						application_keys_decrypts.NewApplication(),
						application_keys_encrypts.NewApplication(),
						1048,
					),
					application_signatures.NewApplication(
						application_votes.NewApplication(
							application_votes_creates.NewApplication(),
							application_votes_validates.NewApplication(),
						),
						application_signs.NewApplication(
							application_signs_creates.NewApplication(),
							application_signs_validates.NewApplication(),
						),
					),
				),
			),
			application_lists.NewApplication(
				application_fetches.NewApplication(),
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

func TestExecute_containsError_returnsError(t *testing.T) {
	firstVariable := "firstVar"
	secondVariable := "secondVar"

	frame := stacks.NewFrameForTests()

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
				application_cryptography_keys.NewApplication(
					application_encryptions.NewApplication(
						application_keys_decrypts.NewApplication(),
						application_keys_encrypts.NewApplication(),
						1048,
					),
					application_signatures.NewApplication(
						application_votes.NewApplication(
							application_votes_creates.NewApplication(),
							application_votes_validates.NewApplication(),
						),
						application_signs.NewApplication(
							application_signs_creates.NewApplication(),
							application_signs_validates.NewApplication(),
						),
					),
				),
			),
			application_lists.NewApplication(
				application_fetches.NewApplication(),
			),
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
}
