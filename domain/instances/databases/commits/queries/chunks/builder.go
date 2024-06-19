package chunks

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        []string
	fingerprint hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        nil,
		fingerprint: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// WithFingerprint adds a fingerprint to the builder
func (app *builder) WithFingerprint(fingerprint hash.Hash) Builder {
	app.fingerprint = fingerprint
	return app
}

// Now builds a new Chunk instance
func (app *builder) Now() (Chunk, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Chunk instance")
	}

	if app.fingerprint == nil {
		return nil, errors.New("the fingerprint hash is mandatory ion order to build a Chunk instance")
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(path),
		app.fingerprint.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createChunk(*pHash, app.path, app.fingerprint), nil
}
