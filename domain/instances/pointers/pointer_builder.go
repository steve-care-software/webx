package pointers

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
	"github.com/steve-care-software/historydb/domain/hash"
)

type pointerBuilder struct {
	hashAdapter hash.Adapter
	path        []string
	isActive    bool
	loader      conditions.Condition
	canceller   conditions.Condition
}

func createPointerBuilder(
	hashAdapter hash.Adapter,
) PointerBuilder {
	out := pointerBuilder{
		hashAdapter: hashAdapter,
		path:        nil,
		isActive:    false,
		loader:      nil,
		canceller:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *pointerBuilder) WithPath(path []string) PointerBuilder {
	app.path = path
	return app
}

// WithLoader adds a loader to the builder
func (app *pointerBuilder) WithLoader(loader conditions.Condition) PointerBuilder {
	app.loader = loader
	return app
}

// WithCanceller adds a canceller to the builder
func (app *pointerBuilder) WithCanceller(canceller conditions.Condition) PointerBuilder {
	app.canceller = canceller
	return app
}

// IsActive flags the builder as active
func (app *pointerBuilder) IsActive() PointerBuilder {
	app.isActive = true
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Pointer instance")
	}

	isActive := "false"
	if app.isActive {
		isActive = "true"
	}

	path := filepath.Join(app.path...)
	data := [][]byte{
		[]byte(path),
		[]byte(isActive),
	}

	if app.loader != nil {
		data = append(data, app.loader.Hash().Bytes())
	}

	if app.canceller != nil {
		data = append(data, app.canceller.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.loader != nil && app.canceller != nil {
		return createPointerWithLoaderAndCanceller(*pHash, app.path, app.isActive, app.loader, app.canceller), nil
	}

	if app.loader != nil {
		return createPointerWithLoader(*pHash, app.path, app.isActive, app.loader), nil
	}

	if app.canceller != nil {
		return createPointerWithCanceller(*pHash, app.path, app.isActive, app.canceller), nil
	}

	return createPointer(*pHash, app.path, app.isActive), nil
}
