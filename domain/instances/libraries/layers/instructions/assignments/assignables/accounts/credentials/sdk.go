package credentials

import "github.com/steve-care-software/datastencil/domain/hash"

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
