package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers"
)

// NewIdentityWithAuthenticatedAndCurrentForTests creates a new identity with authenticated and current for tests
func NewIdentityWithAuthenticatedAndCurrentForTests(all storages.Storages, authenticated switchers.Switchers, current switchers.Switcher) Identity {
	ins, err := NewBuilder().Create().WithAll(all).WithAuthenticated(authenticated).WithCurrent(current).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIdentityWithAuthenticatedForTests creates a new identity with authenticated for tests
func NewIdentityWithAuthenticatedForTests(all storages.Storages, authenticated switchers.Switchers) Identity {
	ins, err := NewBuilder().Create().WithAll(all).WithAuthenticated(authenticated).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIdentityForTests creates a new identity for tests
func NewIdentityForTests(all storages.Storages) Identity {
	ins, err := NewBuilder().Create().WithAll(all).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
