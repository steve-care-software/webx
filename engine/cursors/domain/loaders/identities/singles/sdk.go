package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/profiles"
)

// Adapter represents the single adapter
type Adapter interface {
	ToBytes(ins Single) ([]byte, error)
	ToInstance(data []byte) (Single, error)
}

// Builder represents the single builder
type Builder interface {
	Create() Builder
	WithProfile(profile profiles.Profile) Builder
	WithKeys(keys keys.Keys) Builder
	Now() (Single, error)
}

// Single represents a single identity
type Single interface {
	Profile() profiles.Profile
	Keys() keys.Keys
}
