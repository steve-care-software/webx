package criterias

import (
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an update criteria builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithPassword(password []byte) Builder
	WithSigner(signer signers.Signer) Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	Now() (Criteria, error)
}

// Criteria represents an update criteria
type Criteria interface {
	HasSigner() bool
	Signer() signers.Signer
	HasEncryptor() bool
	Encryptor() encryptors.Encryptor
	HasUsername() bool
	Username() string
	HasPassword() bool
	Password() []byte
}
