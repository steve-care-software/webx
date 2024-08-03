package entities

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

// Adapter represents the entity adapter
type Adapter interface {
	ToBytes(ins Entity) ([]byte, error)
	ToInstance(data []byte) (Entity, error)
}

// Entity represents an entity
type Entity interface {
	Hash() hash.Hash
}
