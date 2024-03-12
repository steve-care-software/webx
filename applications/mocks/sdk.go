package mocks

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
)

// NewAccountRepositoryWithExistsForTests creates a new account repository with exists for tests
func NewAccountRepositoryWithExistsForTests(
	exists map[string]bool,
) accounts.Repository {
	return createAccountRepository(
		exists,
		map[string]accounts.Account{},
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
