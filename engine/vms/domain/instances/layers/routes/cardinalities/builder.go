package cardinalities

import (
	"strconv"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	min         uint
	pMax        *uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		min:         0,
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

// WithMin adds a min to the builder
func (app *builder) WithMin(min uint) Builder {
	app.min = min
	return app
}

// WithMax adds a max to the builder
func (app *builder) WithMax(max uint) Builder {
	app.pMax = &max
	return app
}

// Now builds a new Cardinality instance
func (app *builder) Now() (Cardinality, error) {
	data := [][]byte{
		[]byte(strconv.Itoa(int(app.min))),
	}

	if app.pMax != nil {
		data = append(data, []byte(strconv.Itoa(int(*app.pMax))))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pMax != nil {
		return createCardinalityWithMax(*pHash, app.min, app.pMax), nil
	}

	return createCardinality(*pHash, app.min), nil
}
