package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/profiles"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewSingleBuilder creates a new single builder
func NewSingleBuilder() SingleBuilder {
	return createSingleBuilder()
}

// Adapter represents the single adapter
type Adapter interface {
	ToBytes(ins Single) ([]byte, error)
	ToInstance(data []byte) (Single, error)
}

// Builder represents a singles builder
type Builder interface {
	Create() Builder
	WithList(list []Single) Builder
	Now() (Singles, error)
}

// Singles represents singles
type Singles interface {
	List() []Single
}

// SingleBuilder represents a single builder
type SingleBuilder interface {
	Create() SingleBuilder
	WithProfile(profile profiles.Profile) SingleBuilder
	WithKey(key keys.Key) SingleBuilder
	Now() (Single, error)
}

// Single represents a single identity
type Single interface {
	Profile() profiles.Profile
	Key() keys.Key
}
