package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/storages"
)

type identity struct {
	all           storages.Storages
	authenticated keys.Keys
	current       keys.Key
}

func createIdentity(
	all storages.Storages,
) Identity {
	return createIdentityInternally(all, nil, nil)
}

func createIdentityWithAuthenticated(
	all storages.Storages,
	authenticated keys.Keys,
) Identity {
	return createIdentityInternally(all, authenticated, nil)
}

func createIdentityWithAuthenticatedAndCurrent(
	all storages.Storages,
	authenticated keys.Keys,
	current keys.Key,
) Identity {
	return createIdentityInternally(all, authenticated, current)
}

func createIdentityInternally(
	all storages.Storages,
	authenticated keys.Keys,
	current keys.Key,
) Identity {
	out := identity{
		all:           all,
		authenticated: authenticated,
		current:       current,
	}

	return &out
}

// All returns the identity names
func (obj *identity) All() storages.Storages {
	return obj.all
}

// HasAuthenticated returns true if there is authenticated, false otherwise
func (obj *identity) HasAuthenticated() bool {
	return obj.authenticated != nil
}

// Authenticated returns the authenticated, if any
func (obj *identity) Authenticated() keys.Keys {
	return obj.authenticated
}

// HasCurrent returns true if there is a current identity, false otherwise
func (obj *identity) HasCurrent() bool {
	return obj.current != nil
}

// Current returns the current, if any
func (obj *identity) Current() keys.Key {
	return obj.current
}
