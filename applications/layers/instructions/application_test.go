package instructions

import (
	"testing"

	application_accounts "github.com/steve-care-software/datastencil/applications/layers/instructions/accounts"
	application_accounts_inserts "github.com/steve-care-software/datastencil/applications/layers/instructions/accounts/inserts"
	application_accounts_updates "github.com/steve-care-software/datastencil/applications/layers/instructions/accounts/updates"
	application_assignments "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments"
	application_assignments_assignables "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables"
	application_assignments_assignables_accounts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts"
	application_assignments_assignables_accounts_communications "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications"
	application_assignments_assignables_accounts_communications_signs "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/signs"
	application_assignments_assignables_accounts_communications_votes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/votes"
	application_assignments_assignables_accounts_credentials "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/credentials"
	application_assignments_assignables_accounts_encryptions "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions"
	application_assignments_assignables_accounts_encryptions_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	application_assignments_assignables_accounts_encryptions_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	application_assignments_assignables_accounts_retrieves "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/retrieves"
	application_assignments_assignables_bytes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/bytes"
	application_assignments_assignables_constants "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/constants"
	application_assignments_assignables_cryptography "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography"
	application_assignments_assignables_cryptography_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/decrypts"
	application_assignments_assignables_cryptography_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/encrypts"
	application_assignments_libraries "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries"
	application_assignments_libraries_compilers "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/compilers"
	application_assignments_libraries_databases "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases"
	application_assignments_libraries_databases_repositories "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases/repositories"
	application_assignments_libraries_databases_services "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases/services"
	application_databases "github.com/steve-care-software/datastencil/applications/layers/instructions/databases"
	application_databases_deletes "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/deletes"
	application_databases_inserts "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/inserts"
	application_databases_reverts "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_WithIsStop_Success(t *testing.T) {
	bitRate := 4096

	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithStopForTests(),
	})

	accountRepository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	accountService := mocks.NewAccountServiceForTests(
		true,
		false,
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
	instanceRepository := mocks.NewInstanceRepository(
		23,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	context := uint(45)
	instanceService := mocks.NewInstanceService(
		&context,
		false,
	)

	application := NewApplication(
		application_accounts.NewApplication(
			application_accounts_inserts.NewApplication(
				accountRepository,
				accountService,
				bitRate,
			),
			application_accounts_updates.NewApplication(
				accountRepository,
				accountService,
				bitRate,
			),
			accountService,
		),
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_assignables_accounts.NewApplication(
					application_assignments_assignables_accounts_communications.NewApplication(
						application_assignments_assignables_accounts_communications_signs.NewApplication(),
						application_assignments_assignables_accounts_communications_votes.NewApplication(),
					),
					application_assignments_assignables_accounts_credentials.NewApplication(),
					application_assignments_assignables_accounts_encryptions.NewApplication(
						application_assignments_assignables_accounts_encryptions_decrypts.NewApplication(),
						application_assignments_assignables_accounts_encryptions_encrypts.NewApplication(),
					),
					application_assignments_assignables_accounts_retrieves.NewApplication(
						accountRepository,
					),
					accountRepository,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					),
				),
				application_assignments_libraries.NewApplication(
					application_assignments_libraries_compilers.NewApplication(
						instanceAdapter,
					),
					application_assignments_libraries_databases.NewApplication(
						application_assignments_libraries_databases_repositories.NewApplication(
							instanceRepository,
							skeleton,
						),
						application_assignments_libraries_databases_services.NewApplication(
							instanceService,
						),
					),
				),
			),
		),
		application_databases.NewApplication(
			application_databases_deletes.NewApplication(
				instanceRepository,
				instanceService,
			),
			application_databases_inserts.NewApplication(
				instanceRepository,
				instanceService,
			),
			application_databases_reverts.NewApplication(
				instanceService,
			),
			instanceService,
		),
	)

	stack := stacks.NewFactory().Create()
	stop, _, _, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !stop {
		t.Errorf("the instructions were expected to be stopped")
		return
	}
}
