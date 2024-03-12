package mocks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/criterias"
)

type accountService struct {
	insertWorks bool
	updateWorks bool
	deleteWorks bool
}

func createAccountService(
	insertWorks bool,
	updateWorks bool,
	deleteWorks bool,
) accounts.Service {
	out := accountService{
		insertWorks: insertWorks,
		updateWorks: updateWorks,
		deleteWorks: deleteWorks,
	}

	return &out
}

// Insert inserts an account
func (app *accountService) Insert(account accounts.Account, password []byte) error {
	if app.insertWorks {
		return nil
	}

	return errors.New("insert fails")
}

// Update updates an account
func (app *accountService) Update(credentials credentials.Credentials, criteria criterias.Criteria) error {
	if app.updateWorks {
		return nil
	}

	return errors.New("update fails")
}

// Delete deletes an account
func (app *accountService) Delete(credentials credentials.Credentials) error {
	if app.deleteWorks {
		return nil
	}

	return errors.New("delete fails")
}
