package accounts

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

// Application represents the account application
type Application interface {
	List(password []byte) ([]string, error)
	Retrieve(password []byte, credentials credentials.Credentials) (accounts.Account, error)
	Update(credentials credentials.Credentials, account accounts.Account, criteria accounts.UpdateCriteria) error
	Delete(account accounts.Account, credentials credentials.Credentials) error
	Sign(message []byte, account accounts.Account) (signers.Signature, error)
	Vote(message []byte, ring []signers.PublicKey, account accounts.Account) (signers.Vote, error)
	Encrypt(message []byte, account accounts.Account) ([]byte, error)
	Decrypt(cipher []byte, account accounts.Account) ([]byte, error)
}
