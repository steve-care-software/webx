package executables

import (
	"errors"

	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	local       string
	remote      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		local:       "",
		remote:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithLocal adds a local path to the builder
func (app *builder) WithLocal(local string) Builder {
	app.local = local
	return app
}

// WithRemote adds a remote host to the builder
func (app *builder) WithRemote(remote string) Builder {
	app.remote = remote
	return app
}

// Now builds a new Executable instance
func (app *builder) Now() (Executable, error) {
	data := [][]byte{}
	if app.local != "" {
		data = append(data, []byte("local"))
		data = append(data, []byte(app.local))
	}

	if app.remote != "" {
		data = append(data, []byte("remote"))
		data = append(data, []byte(app.remote))
	}

	if len(data) != 2 {
		return nil, errors.New("the Executable is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.local != "" {
		return createExecutableWithLocal(*pHash, app.local), nil
	}

	return createExecutableWithRemote(*pHash, app.remote), nil
}
