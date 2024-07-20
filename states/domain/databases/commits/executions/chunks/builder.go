package chunks

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        []string
	fingerPrint hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	return &builder{
		hashAdapter: hashAdapter,
		path:        nil,
		fingerPrint: nil,
	}
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the buiolder
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// WithFingerPrint adds a fingerprint to the builder
func (app *builder) WithFingerPrint(fingerPrint hash.Hash) Builder {
	app.fingerPrint = fingerPrint
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

	if app.fingerPrint == nil {
		return nil, errors.New("the fingerPrint is mandatory in order to build a Chunk instance")
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(path),
		app.fingerPrint.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createChunk(
		*pHash,
		app.path,
		app.fingerPrint,
	), nil
}
