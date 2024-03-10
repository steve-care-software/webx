package accounts

import (
	"github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/accounts/criterias"
	account_encryptors "github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the account adapter
type Adapter interface {
	ToBytes(ins Account) ([]byte, error)
	ToInstance(bytes []byte) (Account, error)
}

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithEncryptor(encryptor account_encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	Now() (Account, error)
}

// Account represents the identity account
type Account interface {
	Username() string
	Encryptor() account_encryptors.Encryptor
	Signer() signers.Signer
}

// Repository represents the account repository
type Repository interface {
	List(password []byte) ([]string, error)
	Exists(username string) (bool, error)
	Retrieve(password []byte, credentials credentials.Credentials) (Account, error)
}

// Service represents the account service
type Service interface {
	Insert(account Account, password []byte) error
	Update(credentials credentials.Credentials, criteria criterias.Criteria) error
	Delete(credentials credentials.Credentials) error
}
