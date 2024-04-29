package criterias

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the criteria adapter
type Adapter interface {
	ToBytes(ins Criteria) ([]byte, error)
	ToInstance(bytes []byte) (Criteria, error)
}

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithPassword(password string) Builder
	ChangeSigner() Builder
	ChangeEncryptor() Builder
	Now() (Criteria, error)
}

// Criteria represents an update criteria
type Criteria interface {
	Hash() hash.Hash
	ChangeSigner() bool
	ChangeEncryptor() bool
	HasUsername() bool
	Username() string
	HasPassword() bool
	Password() string
}
