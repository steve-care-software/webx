package mocks

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
)

type accountRepository struct {
	exists map[string]bool
}

func createAccountRepository(
	exists map[string]bool,
) accounts.Repository {
	out := accountRepository{
		exists: exists,
	}

	return &out
}

// List lists the account names
func (app *accountRepository) List(password []byte) ([]string, error) {
	return nil, nil
}

// Exists returns true if exists, false otherwise
func (app *accountRepository) Exists(username string) (bool, error) {
	if value, ok := app.exists[username]; ok {
		return value, nil
	}

	return false, nil
}

// Retrieve retrieves an account
func (app *accountRepository) Retrieve(password []byte, credentials credentials.Credentials) (accounts.Account, error) {
	return nil, nil
}
