package sqllites

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/criterias"
)

type accountService struct {
}

func createAccountService() accounts.Service {
	out := accountService{}

	return &out
}

// Insert inserts an account
func (app *accountService) Insert(account accounts.Account, password []byte) error {
	return nil
}

// Update updates an account
func (app *accountService) Update(credentials credentials.Credentials, criteria criterias.Criteria) error {
	return nil
}

// Delete deletes an account
func (app *accountService) Delete(credentials credentials.Credentials) error {
	return nil
}
