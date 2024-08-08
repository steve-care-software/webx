package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys"
)

// Adapter represents the keys adapter
type Adapter interface {
	ToBytes(ins Single) ([]byte, error)
	ToInstance(data []byte) (Single, error)
}

// SingleBuilder represents a single builder
type Builder interface {
	Create() Builder
	WithKeys(keys keys.Keys) Builder
	Now() (Single, error)
}

// Single represents a single identity
type Single interface {
	Keys() keys.Keys
}
