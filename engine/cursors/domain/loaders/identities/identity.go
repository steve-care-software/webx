package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/storages"
)

type identity struct {
	all           storages.Storages
	authenticated singles.Singles
	current       singles.Single
}

func createIdentity(
	all storages.Storages,
) Identity {
	return createIdentityInternally(all, nil, nil)
}

func createIdentityWithAuthenticated(
	all storages.Storages,
	authenticated singles.Singles,
) Identity {
	return createIdentityInternally(all, authenticated, nil)
}

func createIdentityWithAuthenticatedAndCurrent(
	all storages.Storages,
	authenticated singles.Singles,
	current singles.Single,
) Identity {
	return createIdentityInternally(all, authenticated, current)
}

func createIdentityInternally(
	all storages.Storages,
	authenticated singles.Singles,
	current singles.Single,
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
func (obj *identity) Authenticated() singles.Singles {
	return obj.authenticated
}

// HasCurrent returns true if there is a current identity, false otherwise
func (obj *identity) HasCurrent() bool {
	return obj.current != nil
}

// Current returns the current, if any
func (obj *identity) Current() singles.Single {
	return obj.current
}
