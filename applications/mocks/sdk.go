package mocks

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances"
)

// NewAccountRepositoryWithExistsForTests creates a new account repository with exists for tests
func NewAccountRepositoryWithExistsForTests(
	exists map[string]bool,
) accounts.Repository {
	return createAccountRepository(
		exists,
		map[string]accounts.Account{},
		map[string][]string{},
	)
}

// NewAccountRepositoryWithRetrieveForTests creates a new account repository with retrieve for tests
func NewAccountRepositoryWithRetrieveForTests(
	instances map[string]accounts.Account,
) accounts.Repository {
	exists := map[string]bool{}
	for keyname := range instances {
		exists[keyname] = true
	}

	return createAccountRepository(
		exists,
		instances,
		map[string][]string{},
	)
}

// NewAccountRepositoryWithListForTests creates a new account repository with list for tests
func NewAccountRepositoryWithListForTests(
	list map[string][]string,
) accounts.Repository {
	return createAccountRepository(
		map[string]bool{},
		map[string]accounts.Account{},
		list,
	)
}

// NewAccountServiceForTests creates a new account service for tests
func NewAccountServiceForTests(
	insertWorks bool,
	updateWorks bool,
	deleteWorks bool,
) accounts.Service {
	return createAccountService(
		insertWorks,
		updateWorks,
		deleteWorks,
	)
}

// NewEncryptor creates a new encryptor
func NewEncryptor(
	encrypt map[string]map[string][]byte,
	decrypt map[string]map[string][]byte,
) encryptors.Encryptor {
	return createEncryptor(
		encrypt,
		decrypt,
	)
}

// NewInstanceAdapter creates a new instance adapter
func NewInstanceAdapter(
	toData map[string][]byte,
	toInstance map[string]instances.Instance,
) instances.Adapter {
	return createInstanceAdapter(
		toData,
		toInstance,
	)
}
