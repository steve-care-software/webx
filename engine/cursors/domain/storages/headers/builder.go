package headers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
)

type builder struct {
	identities storages.Storage
}

func createBuilder() Builder {
	out := builder{
		identities: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIdentities add identities to the builder
func (app *builder) WithIdentities(identities storages.Storage) Builder {
	app.identities = identities
	return app
}

// Now builds a new Header instance
func (app *builder) Now() (Header, error) {
	if app.identities != nil {
		return createHeaderWithIdentities(app.identities), nil
	}

	return nil, errors.New("the Header is invalid")
}
