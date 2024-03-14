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
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts"
	account_inserts "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables"
	assignable_bytes "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/inserts"
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
	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsStop() {
		t.Errorf("the instructions were expected to be stopped")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithRaisedError_Success(t *testing.T) {
	bitRate := 4096

	raisedError := uint(45)
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithRaiseErrorForTests(raisedError),
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
	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsFailure() {
		t.Errorf("the instructions were expected to be a failure")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithAccount_Succeeds_Success(t *testing.T) {
	bitRate := 4096
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

	userVar := "username"
	username := "myUsername"
	passVar := "password"
	password := "myPassword"
	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(userVar, stacks.NewAssignableWithBytesForTests([]byte(username))),
					stacks.NewAssignmentForTests(passVar, stacks.NewAssignableWithBytesForTests([]byte(password))),
				}),
			),
		}),
	)

	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithAccountForTests(
			accounts.NewAccountWithInsertForTests(
				account_inserts.NewInsertForTests(userVar, passVar),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption != nil {
		t.Errorf("the interruption was expected to NOT be valid")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithAccount_Fails_ReturnsFailure(t *testing.T) {
	bitRate := 4096
	accountRepository := mocks.NewAccountRepositoryWithExistsForTests(map[string]bool{})
	accountService := mocks.NewAccountServiceForTests(
		false, // insert does not work
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

	userVar := "username"
	username := "myUsername"
	passVar := "password"
	password := "myPassword"
	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(userVar, stacks.NewAssignableWithBytesForTests([]byte(username))),
					stacks.NewAssignmentForTests(passVar, stacks.NewAssignableWithBytesForTests([]byte(password))),
				}),
			),
		}),
	)

	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithAccountForTests(
			accounts.NewAccountWithInsertForTests(
				account_inserts.NewInsertForTests(userVar, passVar),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsFailure() {
		t.Errorf("the instructions were expected to be a failure")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}

	failure := retInterruption.Failure()
	if !failure.HasMessage() {
		t.Errorf("the failure was expected to contain a message")
		return
	}
}

func TestExecute_WithDatabase_Succeeds_Success(t *testing.T) {
	bitRate := 4096
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

	contextVar := "myContext"
	pathVar := "myPath"
	path := []string{
		"this",
		"is",
		"a",
		"path",
	}

	instanceVar := "myInstance"
	instance := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)
	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						contextVar,
						stacks.NewAssignableWithUnsignedIntForTests(context),
					),
					stacks.NewAssignmentForTests(
						instanceVar,
						stacks.NewAssignableWithInstanceForTests(instance),
					),
					stacks.NewAssignmentForTests(
						pathVar,
						stacks.NewAssignableWithStringListForTests(path),
					),
				}),
			),
		}),
	)

	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithDatabaseForTests(
			databases.NewDatabaseWithInsertForTests(
				inserts.NewInsertForTests(contextVar, instanceVar, pathVar),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption != nil {
		t.Errorf("the interruption was expected to NOT be valid")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithDatabase_Fails_ReturnsFailure(t *testing.T) {
	bitRate := 4096
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

	contextVar := "myContext"
	pathVar := "myPath"
	path := []string{
		"this",
		"is",
		"a",
		"path",
	}

	instanceVar := "myInstance"
	instance := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)
	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						instanceVar,
						stacks.NewAssignableWithInstanceForTests(instance),
					),
					stacks.NewAssignmentForTests(
						pathVar,
						stacks.NewAssignableWithStringListForTests(path),
					),
				}),
			),
		}),
	)

	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithDatabaseForTests(
			databases.NewDatabaseWithInsertForTests(
				inserts.NewInsertForTests(contextVar, instanceVar, pathVar),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsFailure() {
		t.Errorf("the instructions were expected to be a failure")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}

	failure := retInterruption.Failure()
	if !failure.HasMessage() {
		t.Errorf("the failure was expected to contain a message")
		return
	}
}

func TestExecute_WithAssignment_Succeeds_Success(t *testing.T) {
	bitRate := 4096
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

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
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
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption != nil {
		t.Errorf("the interruption was expected to NOT be valid")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}

	retStackAssignable, err := retStack.Head().Fetch(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !retStackAssignable.IsBytes() {
		t.Errorf("the returned assignable was expected to contain bytes")
		return
	}
}

func TestExecute_WithAssignment_Fails_ReturnsFailure(t *testing.T) {
	bitRate := 4096
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

	firstVariable := "firstVar"
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						secondVariable,
						stacks.NewAssignableWithBytesForTests(
							secondValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsFailure() {
		t.Errorf("the instructions were expected to be a failure")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}

	failure := retInterruption.Failure()
	if !failure.HasMessage() {
		t.Errorf("the failure was expected to contain a message")
		return
	}
}

func TestExecute_WithCondition_ConditionIsFalse_WithAssignment_Succeeds_Success(t *testing.T) {
	bitRate := 4096
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

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	conditionVar := "myCondition"
	conditionValue := false

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
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
					stacks.NewAssignmentForTests(
						conditionVar,
						stacks.NewAssignableWithBoolForTests(
							conditionValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithConditionForTests(
			instructions.NewConditionForTest(
				conditionVar,
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithStopForTests(),
				}),
			),
		),
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption != nil {
		t.Errorf("the interruption was expected to NOT be valid")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}

	retStackAssignable, err := retStack.Head().Fetch(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !retStackAssignable.IsBytes() {
		t.Errorf("the returned assignable was expected to contain bytes")
		return
	}
}

func TestExecute_WithCondition_ConditionIsTrue_WithAssignment_ExecutesConditionInstructions_Stops_ReturnsInterruption(t *testing.T) {
	bitRate := 4096
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

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	conditionVar := "myCondition"
	conditionValue := true

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
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
					stacks.NewAssignmentForTests(
						conditionVar,
						stacks.NewAssignableWithBoolForTests(
							conditionValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithConditionForTests(
			instructions.NewConditionForTest(
				conditionVar,
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithStopForTests(),
				}),
			),
		),
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsStop() {
		t.Errorf("the instructions were expected to be stopped")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithCondition_ConditionIsTrue_WithAssignment_ExecutesConditionInstructions_Failure_ReturnsInterruption(t *testing.T) {
	bitRate := 4096
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

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	conditionVar := "myCondition"
	conditionValue := true

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
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
					stacks.NewAssignmentForTests(
						conditionVar,
						stacks.NewAssignableWithBoolForTests(
							conditionValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithConditionForTests(
			instructions.NewConditionForTest(
				conditionVar,
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithRaiseErrorForTests(45),
				}),
			),
		),
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsFailure() {
		t.Errorf("the instructions were expected to be failure")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithCondition_ConditionIsTrue_WithAssignment_ExecutesConditionInstructions_Succeeds_Success(t *testing.T) {
	bitRate := 4096
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

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	conditionVar := "myCondition"
	conditionValue := true

	contextVar := "myContext"
	pathVar := "myPath"
	path := []string{
		"this",
		"is",
		"a",
		"path",
	}

	instanceVar := "myInstance"
	instance := inserts.NewInsertForTests(contextVar, instanceVar, pathVar)

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
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
					stacks.NewAssignmentForTests(
						conditionVar,
						stacks.NewAssignableWithBoolForTests(
							conditionValue,
						),
					),
					stacks.NewAssignmentForTests(
						contextVar,
						stacks.NewAssignableWithUnsignedIntForTests(context),
					),
					stacks.NewAssignmentForTests(
						instanceVar,
						stacks.NewAssignableWithInstanceForTests(instance),
					),
					stacks.NewAssignmentForTests(
						pathVar,
						stacks.NewAssignableWithStringListForTests(path),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithConditionForTests(
			instructions.NewConditionForTest(
				conditionVar,
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithDatabaseForTests(
						databases.NewDatabaseWithInsertForTests(
							inserts.NewInsertForTests(contextVar, instanceVar, pathVar),
						),
					),
				}),
			),
		),
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption != nil {
		t.Errorf("the interruption was expected to NOT be valid")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}
