package instances

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
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
