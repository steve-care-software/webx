package resources

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	code           uint
	isRaiseInLayer bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		code:           0,
		isRaiseInLayer: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCode adds a code to the builder
func (app *builder) WithCode(code uint) Builder {
	app.code = code
	return app
}

// IsRaisedInLayer flags the builder as isRaisedInLayer
func (app *builder) IsRaisedInLayer() Builder {
	app.isRaiseInLayer = true
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	if app.code <= 0 {
		return nil, errors.New("the code is mandatory in order to build a Resource instance")
	}

	isRaisedInLayer := "false"
	if app.isRaiseInLayer {
		isRaisedInLayer = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.code))),
		[]byte(isRaisedInLayer),
	})

	if err != nil {
		return nil, err
	}

	return createResource(*pHash, app.code, app.isRaiseInLayer), nil
}
