package inserts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	bytes       []byte
	whitelist   []hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		bytes:       nil,
		whitelist:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// WithWhitelist adds a whitelist to the builder
func (app *builder) WithWhitelist(whitelist []hash.Hash) Builder {
	app.whitelist = whitelist
	return app
}

// Now builds a new Insert instance
func (app *builder) Now() (Insert, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Insert instance")
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes is mandatory in order to build an Insert instance")
	}

	if app.whitelist != nil && len(app.whitelist) <= 0 {
		app.whitelist = nil
	}

	if app.whitelist == nil {
		return nil, errors.New("the whitelist is mandatory in order to build an Insert instance")
	}

	data := [][]byte{
		[]byte(app.name),
		app.bytes,
	}

	for _, oneHash := range app.whitelist {
		data = append(data, oneHash.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createInsert(
		*pHash,
		app.name,
		app.bytes,
		app.whitelist,
	), nil
}
