package criterias

import "github.com/steve-care-software/datastencil/domain/hash"

type builder struct {
	hashAdapter     hash.Adapter
	changeSigner    bool
	changeEncryptor bool
	username        string
	password        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:     hashAdapter,
		changeSigner:    false,
		changeEncryptor: false,
		username:        "",
		password:        "",
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

// ChangeSigner flags the builder as change signer
func (app *builder) ChangeSigner() Builder {
	app.changeSigner = true
	return app
}

// ChangeEncryptor flags the builder as change encryptor
func (app *builder) ChangeEncryptor() Builder {
	app.changeEncryptor = true
	return app
}

// Now builds a new Criteria instance
func (app *builder) Now() (Criteria, error) {
	data := [][]byte{}
	if app.changeSigner {
		data = append(data, []byte("changeSigner"))
	}

	if app.changeEncryptor {
		data = append(data, []byte("changeEncryptor"))
	}

	if app.username != "" {
		data = append(data, []byte(app.username))
	}

	if app.password != "" {
		data = append(data, []byte(app.password))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createCriteria(
		*pHash,
		app.changeSigner,
		app.changeEncryptor,
		app.username,
		app.password,
	), nil
}
