package inserts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	username    string
	password    string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		username:    "",
		password:    "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithUsername adds a username to the builder
func (app *builder) WithUsername(username string) Builder {
	app.username = username
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) Builder {
	app.password = password
	return app
}

// Now builds a new Insert instance
func (app *builder) Now() (Insert, error) {
	if app.username == "" {
		return nil, errors.New("the username is mandatory in order to build an Insert instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build an Insert instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.username),
		[]byte(app.password),
	})

	if err != nil {
		return nil, err
	}

	return createInsert(*pHash,
		app.username,
		app.password,
	), nil
}
