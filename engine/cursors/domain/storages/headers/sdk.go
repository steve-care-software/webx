package headers

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/pointers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the header adapter
type Adapter interface {
	ToBytes(ins Header) ([]byte, error)
	ToInstance(data []byte) (Header, error)
}

// Builder represents the header builder
type Builder interface {
	Create() Builder
	WithIdentities(identities pointers.Pointer) Builder
	Now() (Header, error)
}

// Header represents an header
type Header interface {
	HasIdentities() bool
	Identities() pointers.Pointer
}
