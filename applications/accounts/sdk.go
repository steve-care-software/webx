package accounts

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
)

// Application represents the account application
type Application interface {
	List(password []byte) ([]string, error)
	Retrieve(password []byte, credentials credentials.Credentials) (accounts.Account, error)
	Update(credentials credentials.Credentials, account accounts.Account, criteria accounts.UpdateCriteria) error
	Delete(credentials credentials.Credentials, account accounts.Account) error
}
