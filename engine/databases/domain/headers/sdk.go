package headers

import (
	"github.com/steve-care-software/webx/engine/databases/domain/headers/states"
)

// Adaptetr represents an header adapter
type Adaptetr interface {
	ToInstance(bytes []byte) (Header, error)
	ToBytes(ins Header) ([]byte, error)
}

// Builder represents an header builder
type Builder interface {
	Create() Builder
	WithLength(length uint64) Builder
	WithStates(states states.States) Builder
	Now() (Header, error)
}

// Header represents an header
type Header interface {
	Length() uint64
	States() states.States
}
