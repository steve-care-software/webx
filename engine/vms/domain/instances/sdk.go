package instances

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

// Adapter represents the instance adapter
type Adapter interface {
	ToBytes(ins Instance) ([]byte, error)
	ToInstance(data []byte) (Instance, error)
}

// Instance represents an instance
type Instance interface {
	Hash() hash.Hash
}

// Repository represents an instance reposiory
type Repository interface {
	Amount(kind string) (*uint, error)
	List(kind string, index uint, amount uint) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) ([]byte, error)
}

// Service represents an instance service
type Service interface {
	Insert(kind string, hash hash.Hash, data []byte) error
	Delete(kind string, hash hash.Hash) error
}
