package mocks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
)

type accountRepository struct {
	exists   map[string]bool
	retrieve map[string]accounts.Account
	list     map[string][]string
}

func createAccountRepository(
	exists map[string]bool,
	retrieve map[string]accounts.Account,
	list map[string][]string,
) accounts.Repository {
	out := accountRepository{
		exists:   exists,
		retrieve: retrieve,
		list:     list,
	}

	return &out
}

// List lists the account names
func (app *accountRepository) List(password []byte) ([]string, error) {
	if list, ok := app.list[string(password)]; ok {
		return list, nil
	}

	return nil, errors.New("there is no list attached to the provided password")
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
	username := credentials.Username()
	if ins, ok := app.retrieve[username]; ok {
		return ins, nil
	}

	return nil, errors.New("there is no account attached to the provided username")
}
