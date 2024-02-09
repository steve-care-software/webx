package results

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/identity/domain/hash"
)

type failureBuilder struct {
	hashAdapter     hash.Adapter
	code            uint
	isRaisedInLayer bool
}

func createFailureBuilder(
	hashAdapter hash.Adapter,
) FailureBuilder {
	out := failureBuilder{
		hashAdapter:     hashAdapter,
		code:            0,
		isRaisedInLayer: false,
	}

	return &out
}

// Create initializes the builder
func (app *failureBuilder) Create() FailureBuilder {
	return createFailureBuilder(
		app.hashAdapter,
	)
}

// WithCode adds a code to the builder
func (app *failureBuilder) WithCode(code uint) FailureBuilder {
	app.code = code
	return app
}

// IsRaisedInLayer flags the builder as isRaisedInLayer
func (app *failureBuilder) IsRaisedInLayer() FailureBuilder {
	app.isRaisedInLayer = true
	return app
}

// Now builds a new Failure instance
func (app *failureBuilder) Now() (Failure, error) {
	if app.code <= 0 {
		return nil, errors.New("the code is mandatory in order to build a Failure instance")
	}

	isRaisedInLayer := "false"
	if app.isRaisedInLayer {
		isRaisedInLayer = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.code))),
		[]byte(isRaisedInLayer),
	})

	if err != nil {
		return nil, err
	}

	return createFailure(*pHash, app.code, app.isRaisedInLayer), nil
}
