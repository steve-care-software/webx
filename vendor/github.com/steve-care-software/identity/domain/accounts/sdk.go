package accounts

import (
	account_encryptors "github.com/steve-care-software/identity/domain/accounts/encryptors"
	"github.com/steve-care-software/identity/domain/accounts/signers"
	"github.com/steve-care-software/identity/domain/credentials"
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

// UpdateCriteriaBuilder represents an update criteria builder
type UpdateCriteriaBuilder interface {
	Create() UpdateCriteriaBuilder
	WithUsername(username string) UpdateCriteriaBuilder
	WithPassword(password []byte) UpdateCriteriaBuilder
	ChangeSigner() UpdateCriteriaBuilder
	ChangeEncryptor() UpdateCriteriaBuilder
	Now() (UpdateCriteria, error)
}

// UpdateCriteria represents an update criteria
type UpdateCriteria interface {
	ChangeSigner() bool
	ChangeEncryptor() bool
	HasUsername() bool
	Username() string
	HasPassword() bool
	Password() []byte
}

// Repository represents the account repository
type Repository interface {
	List() ([]string, error)
	Exists(username string) (bool, error)
	Retrieve(credentials credentials.Credentials) (Account, error)
}

// Service represents the account service
type Service interface {
	Insert(account Account, password []byte) error
	Update(credentials credentials.Credentials, criteria UpdateCriteria) error
	Delete(credentials credentials.Credentials) error
}
