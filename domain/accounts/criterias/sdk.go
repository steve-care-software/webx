package criterias

import (
	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

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
	ChangeSigner() bool
	ChangeEncryptor() bool
	HasUsername() bool
	Username() string
	HasPassword() bool
	Password() []byte
}
