package credentials

import "errors"

type builder struct {
	username string
	password []byte
}

func createBuilder() Builder {
	out := builder{
		username: "",
		password: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithUsername adds a username to the builder
func (app *builder) WithUsername(username string) Builder {
	app.username = username
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password []byte) Builder {
	app.password = password
	return app
}

// Now builds a new Credentials instance
func (app *builder) Now() (Credentials, error) {
	if app.username == "" {
		return nil, errors.New("the username is mandatory in order to build a Credentials instance")
	}

	if app.password != nil && len(app.password) <= 0 {
		app.password = nil
	}

	if app.password == nil {
		return nil, errors.New("the password is mandatory in order to build a Credentials instance")
	}

	return createCredentials(app.username, app.password), nil
}
