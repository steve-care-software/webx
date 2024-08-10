package identities

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/storages"
)

type builder struct {
	all           storages.Storages
	authenticated keys.Keys
	current       keys.Key
}

func createBuilder() Builder {
	out := builder{
		all:           nil,
		authenticated: nil,
		current:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithAll add storages to the builder
func (app *builder) WithAll(all storages.Storages) Builder {
	app.all = all
	return app
}

// WithAuthenticated add authenticated keys to the builder
func (app *builder) WithAuthenticated(authenticated keys.Keys) Builder {
	app.authenticated = authenticated
	return app
}

// WithCurrent adds the current key to the builder
func (app *builder) WithCurrent(current keys.Key) Builder {
	app.current = current
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.all == nil {
		return nil, errors.New("the storage identities are mandatory in order to build an Identity instance")
	}

	if app.authenticated != nil && app.current != nil {
		return createIdentityWithAuthenticatedAndCurrent(app.all, app.authenticated, app.current), nil
	}

	if app.authenticated != nil {
		return createIdentityWithAuthenticated(app.all, app.authenticated), nil
	}

	return createIdentity(app.all), nil
}
