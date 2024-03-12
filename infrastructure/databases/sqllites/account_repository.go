package sqllites

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
)

type accountRepository struct {
}

func createAccountRepository() accounts.Repository {
	out := accountRepository{}

	return &out
}

// List lists the account names
func (app *accountRepository) List(password []byte) ([]string, error) {
	return nil, nil
}

// Exists returns true if exists, false otherwise
func (app *accountRepository) Exists(username string) (bool, error) {
	return false, nil
}

// Retrieve retrieves an account
func (app *accountRepository) Retrieve(password []byte, credentials credentials.Credentials) (accounts.Account, error) {
	return nil, nil
}
