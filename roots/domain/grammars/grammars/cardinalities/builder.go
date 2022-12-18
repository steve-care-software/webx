package cardinalities

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pMin        *uint
	pMax        *uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pMin:        nil,
		pMax:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithMin adds a minimum to the builder
func (app *builder) WithMin(min uint) Builder {
	app.pMin = &min
	return app
}

// WithMax adds a maximum to the builder
func (app *builder) WithMax(max uint) Builder {
	app.pMax = &max
	return app
}

// Now builds a new Cardinality instance
func (app *builder) Now() (Cardinality, error) {
	if app.pMin == nil {
		return nil, errors.New("the minimum is mandatory in order to build a Cardinality instance")
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(*app.pMin))),
	}

	if app.pMax != nil {
		data = append(data, []byte(strconv.Itoa(int(*app.pMax))))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pMax != nil {
		return createCardinalityWithMax(*pHash, *app.pMin, app.pMax), nil
	}

	return createCardinality(*pHash, *app.pMin), nil
}
