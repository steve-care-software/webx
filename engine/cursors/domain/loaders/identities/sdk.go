package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/profiles"
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
	Now() (Identitys, error)
}

// Identitys represents identitys
type Identitys interface {
	List() []Identity
}

// IdentityBuilder represents a identity builder
type IdentityBuilder interface {
	Create() IdentityBuilder
	WithProfile(profile profiles.Profile) IdentityBuilder
	WithKey(key keys.Key) IdentityBuilder
	WithNamespaces(namespaces namespaces.Namespaces) IdentityBuilder
	Now() (Identity, error)
}

// Identity represents a identity identity
type Identity interface {
	Profile() profiles.Profile
	Key() keys.Key
	HasNamespaces() bool
	Namespaces() namespaces.Namespaces
}
