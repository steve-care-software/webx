package retrieves

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	password    string
	credentials string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		password:    "",
		credentials: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) Builder {
	app.password = password
	return app
}

// WithCredentials add credentials to the builder
func (app *builder) WithCredentials(credentials string) Builder {
	app.credentials = credentials
	return app
}

// Now builds a new Retrieve instance
func (app *builder) Now() (Retrieve, error) {
	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build a Retrieve instance")
	}

	if app.credentials == "" {
		return nil, errors.New("the credentials is mandatory in order to build a Retrieve instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.password),
		[]byte(app.credentials),
	})

	if err != nil {
		return nil, err
	}

	return createRetrieve(*pHash, app.password, app.credentials), nil
}
