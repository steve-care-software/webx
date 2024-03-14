package assignments

import (
	"bytes"
	"testing"

	application_assignables "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts"
	application_execution_communications "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/communications"
	application_signs "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/communications/signs"
	application_votes "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/communications/votes"
	application_execution_credentials "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/credentials"
	application_execution_encryptions "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/encryptions"
	application_accounts_decrypts "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	application_accounts_encrypts "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	application_execution_retrieves "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/retrieves"
	application_bytes "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/bytes"
	application_constants "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/constants"
	application_cryptography "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/cryptography"
	application_decrypts "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/cryptography/decrypts"
	application_encrypts "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/cryptography/encrypts"
	application_libraries "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries"
	application_compilers "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries/compilers"
	application_databases "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries/databases"
	application_repositories "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries/databases/repositories"
	application_services "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries/databases/services"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables"
	assignable_bytes "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/bytes"
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

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})
	instanceRepository := mocks.NewInstanceRepository(
		23,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	context := uint(33)
	instanceService := mocks.NewInstanceService(
		&context,
		false,
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	skeleton := mocks.NewSkeleton()

	application := NewApplication(
		application_assignables.NewApplication(
			accounts.NewApplication(
				application_execution_communications.NewApplication(
					application_signs.NewApplication(),
					application_votes.NewApplication(),
				),
				application_execution_credentials.NewApplication(),
				application_execution_encryptions.NewApplication(
					application_accounts_decrypts.NewApplication(),
					application_accounts_encrypts.NewApplication(),
				),
				application_execution_retrieves.NewApplication(
					accountRepository,
				),
				accountRepository,
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
			),
			application_libraries.NewApplication(
				application_compilers.NewApplication(
					instanceAdapter,
				),
				application_databases.NewApplication(
					application_repositories.NewApplication(
						instanceRepository,
						skeleton,
					),
					application_services.NewApplication(
						instanceService,
					),
				),
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

func TestExecute_isMistake_returnsError(t *testing.T) {
	queryVar := "queryVar"
	frame := stacks.NewFrameForTests()

	name := "myName"
	instruction := assignments.NewAssignmentForTests(
		name,
		assignables.NewAssignableWithQueryForTests(
			queryVar,
		),
	)

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})
	instanceRepository := mocks.NewInstanceRepository(
		23,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	context := uint(33)
	instanceService := mocks.NewInstanceService(
		&context,
		false,
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	skeleton := mocks.NewSkeleton()

	application := NewApplication(
		application_assignables.NewApplication(
			accounts.NewApplication(
				application_execution_communications.NewApplication(
					application_signs.NewApplication(),
					application_votes.NewApplication(),
				),
				application_execution_credentials.NewApplication(),
				application_execution_encryptions.NewApplication(
					application_accounts_decrypts.NewApplication(),
					application_accounts_encrypts.NewApplication(),
				),
				application_execution_retrieves.NewApplication(
					accountRepository,
				),
				accountRepository,
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
			),
			application_libraries.NewApplication(
				application_compilers.NewApplication(
					instanceAdapter,
				),
				application_databases.NewApplication(
					application_repositories.NewApplication(
						instanceRepository,
						skeleton,
					),
					application_services.NewApplication(
						instanceService,
					),
				),
			),
		),
	)

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
	if code != failures.CouldNotFetchQueryFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchQueryFromFrame, code)
		return
	}
}
