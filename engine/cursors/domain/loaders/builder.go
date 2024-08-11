package loaders

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities"

type builder struct {
	identity identities.Identity
}

func createBuilder() Builder {
	out := builder{
		identity: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIdentity adds an identity to the builder
func (app *builder) WithIdentity(identity identities.Identity) Builder {
	app.identity = identity
	return app
}

// Now builds a new Loader instance
func (app *builder) Now() (Loader, error) {
	if app.identity != nil {
		return createLoaderWithIdentity(app.identity), nil
	}

	return createLoader(), nil
}
