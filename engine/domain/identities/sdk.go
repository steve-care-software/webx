package identities

import (
	"github.com/steve-care-software/webx/engine/domain/identities/keys"
	"github.com/steve-care-software/webx/engine/domain/identities/profiles"
)

// Adapter represents the identity adapter
type Adapter interface {
	ToBytes(ins Identity) ([]byte, error)
	ToInstance(data []byte) (Identity, error)
}

// Builder represents a identitys builder
type Builder interface {
	Create() Builder
	WithList(list []Identity) Builder
	Now() (Identities, error)
}

// Identities represents identitys
type Identities interface {
	List() []Identity
}

// IdentityBuilder represents a identity builder
type IdentityBuilder interface {
	Create() IdentityBuilder
	WithProfile(profile profiles.Profile) IdentityBuilder
	WithKey(key keys.Key) IdentityBuilder
	Now() (Identity, error)
}

// Identity represents a identity identity
type Identity interface {
	Profile() profiles.Profile
	Key() keys.Key
}
