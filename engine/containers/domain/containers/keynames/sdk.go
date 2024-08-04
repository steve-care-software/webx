package keynames

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

// Builder represents a keyname builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	Now() (Keyname, error)
}

// Keyname represents a keyname
type Keyname interface {
	Hash() hash.Hash
	Name() string
}
