package inserts

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Builder represents an insert builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithBytes(bytes []byte) Builder
	WithWhitelist(whitelist []hash.Hash) Builder
	Now() (Insert, error)
}

// Insert represents an insert
type Insert interface {
	Hash() hash.Hash
	Name() string
	Bytes() []byte
	Whitelist() []hash.Hash
}
