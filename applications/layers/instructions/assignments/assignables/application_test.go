package assignables

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts"
	application_execution_communications "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications"
	application_signs "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/signs"
	application_votes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/votes"
	application_execution_credentials "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/credentials"
	application_execution_encryptions "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions"
	application_accounts_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	application_accounts_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	application_execution_retrieves "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/retrieves"
	application_bytes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/bytes"
	application_compilers "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/compilers"
	application_constants "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/constants"
	application_cryptography "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography"
	application_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/decrypts"
	application_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/encrypts"
	application_databases "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases"
	application_repositories "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/repositories"
	application_services "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/services"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables"
	assignables_accounts "github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts"
	assignable_bytes "github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_withBytes_Success(t *testing.T) {
	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	context := uint(45)

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

	instruction := assignables.NewAssignableWithBytesForTests(
		assignable_bytes.NewBytesWithJoinForTests([]string{
			firstVariable,
			secondVariable,
		}),
	)

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	skeleton := mocks.NewSkeleton()

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

func TestExecute_withConstant_Success(t *testing.T) {
	context := uint(45)

	instruction := assignables.NewAssignableWithConstantForTests(
		constants.NewConstantWithBoolForTests(true),
	)

	frame := stacks.NewFrameForTests()

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	skeleton := mocks.NewSkeleton()

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain bool")
		return
	}

	pValue := retAssignable.Bool()
	if !*pValue {
		t.Errorf("the value was expected to be true")
		return
	}
}

func TestExecute_withCryptography_Success(t *testing.T) {
	cipherVar := "cipherVar"
	cipher := []byte("this is a cipher")
	passwordVar := "passVar"
	password := []byte("this is a password")
	message := []byte("this is some message")

	context := uint(45)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests(
					cipher,
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

	instruction := assignables.NewAssignableWithCryptographyForTests(
		cryptography.NewCryptographyWithDecryptForTests(
			decrypts.NewDecryptForTests(cipherVar, passwordVar),
		),
	)

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{
			string(cipher): map[string][]byte{
				string(password): message,
			},
		},
	)

	skeleton := mocks.NewSkeleton()

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retMessage := retAssignable.Bytes()
	if !bytes.Equal(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestExecute_withAccount_Success(t *testing.T) {
	usernameList := []string{
		"first",
		"second",
	}

	passwordVar := "passwordVar"
	password := "myPassword"

	context := uint(45)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					[]byte(password),
				),
			),
		}),
	)

	instruction := assignables.NewAssignableWithAccountForTests(
		assignables_accounts.NewAccountWithListForTests(passwordVar),
	)

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{
		password: usernameList,
	})

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	skeleton := mocks.NewSkeleton()

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

	if !retAssignable.IsStringList() {
		t.Errorf("the assignable was expected to contain a string list")
		return
	}

	retStringList := retAssignable.StringList()
	if !reflect.DeepEqual(usernameList, retStringList) {
		t.Errorf("the returned usernames list is invalid")
		return
	}
}

func TestExecute_withQuery_queryExistsInFrame_Success(t *testing.T) {
	queryVar := "queryVar"
	query := mocks.NewQuery()

	context := uint(45)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				queryVar,
				stacks.NewAssignableWithQueryForTests(
					query,
				),
			),
		}),
	)

	instruction := assignables.NewAssignableWithQueryForTests(
		queryVar,
	)

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	skeleton := mocks.NewSkeleton()

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

	if !retAssignable.IsQuery() {
		t.Errorf("the assignable was expected to contain a query")
		return
	}

	retQuery := retAssignable.Query()
	if !reflect.DeepEqual(query, retQuery) {
		t.Errorf("the returned query is invalid")
		return
	}
}

func TestExecute_withQuery_queryDoesNotExistsInFrame_returnsError(t *testing.T) {
	queryVar := "queryVar"
	frame := stacks.NewFrameForTests()

	context := uint(45)

	instruction := assignables.NewAssignableWithQueryForTests(
		queryVar,
	)

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	skeleton := mocks.NewSkeleton()

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

func TestExecute_WithCompiler_Success(t *testing.T) {
	compileVar := "compileVar"
	compile := []byte("this is some data")
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	context := uint(45)
	skeleton := mocks.NewSkeleton()

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				compileVar,
				stacks.NewAssignableWithBytesForTests(
					compile,
				),
			),
		}),
	)

	instruction := assignables.NewAssignableWithCompilerForTests(
		compilers.NewCompilerWithCompileForTests(compileVar),
	)

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{
					string(compile): compiledInstance,
				},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

	if !retAssignable.IsInstance() {
		t.Errorf("the assignable was expected to contain an instance")
		return
	}

	retInstance := retAssignable.Instance()
	if !reflect.DeepEqual(compiledInstance.(instances.Instance), retInstance) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestExecute_WithDatabase_Success(t *testing.T) {
	context := uint(45)
	skeleton := mocks.NewSkeleton()

	frame := stacks.NewFrameForTests()
	instruction := assignables.NewAssignableWithDatabaseForTests(
		databases.NewDatabaseWithRepositoryForTests(
			repositories.NewRepositoryWithSkeletonForTests(),
		),
	)

	accountRepository := mocks.NewAccountRepositoryWithListForTests(map[string][]string{})

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

	if !retAssignable.IsSkeleton() {
		t.Errorf("the assignable was expected to contain a skeleton")
		return
	}

	retSkeleton := retAssignable.Skeleton()
	if !reflect.DeepEqual(skeleton, retSkeleton) {
		t.Errorf("the returned skeleton is invalid")
		return
	}
}
