package credentials

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the credentials adapter
type Adapter interface {
	ToBytes(ins Credentials) ([]byte, error)
	ToInstance(bytes []byte) (Credentials, error)
}

// Builder represents a credentials builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithPassword(password string) Builder
	Now() (Credentials, error)
}

// Credentials represents a credentials
type Credentials interface {
	Hash() hash.Hash
	Username() string
	Password() string
}
