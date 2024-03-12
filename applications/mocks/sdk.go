package mocks

import "github.com/steve-care-software/datastencil/domain/accounts"

// NewAccountRepositoryWithExistsForTests creates a new account repository with exists for tests
func NewAccountRepositoryWithExistsForTests(
	exists map[string]bool,
) accounts.Repository {
	return createAccountRepository(exists)
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
