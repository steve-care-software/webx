package inserts

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the insert adapter
type Adapter interface {
	ToBytes(ins Insert) ([]byte, error)
	ToInstance(bytes []byte) (Insert, error)
}

// Builder represents an insert builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithPassword(password string) Builder
	Now() (Insert, error)
}

// Insert represents an insert
type Insert interface {
	Hash() hash.Hash
	Username() string
	Password() string
}
