package headers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/hashes/domain/pointers"
)

type builder struct {
	identities pointers.Pointer
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
func (app *builder) WithIdentities(identities pointers.Pointer) Builder {
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
