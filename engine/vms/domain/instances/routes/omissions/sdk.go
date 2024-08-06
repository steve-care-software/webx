package omissions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
)

// NewBuilder creates a new omission builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the omission adapter
type Adapter interface {
	ToBytes(ins Omission) ([]byte, error)
	ToInstance(bytes []byte) (Omission, error)
}

// Builder represents the omission builder
type Builder interface {
	Create() Builder
	WithPrefix(prefix elements.Element) Builder
	WithSuffix(suffix elements.Element) Builder
	Now() (Omission, error)
}

// Omission represents an omission
type Omission interface {
	Hash() hash.Hash
	HasPrefix() bool
	Prefix() elements.Element
	HasSuffix() bool
	Suffix() elements.Element
}
