package services

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	isBegin     bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isBegin:     false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// IsBegin flags the builder as begin
func (app *builder) IsBegin() Builder {
	app.isBegin = true
	return app
}

// Now builds a new Service instance
func (app *builder) Now() (Service, error) {
	data := [][]byte{}
	if app.isBegin {
		data = append(data, []byte("isBegin"))
	}

	if len(data) != 1 {
		return nil, errors.New("the Service is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createServiceWithBegin(*pHash), nil
}
