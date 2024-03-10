package results

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type failureBuilder struct {
	hashAdapter     hash.Adapter
	pIndex          *uint
	pCode           *uint
	isRaisedInLayer bool
}

func createFailureBuilder(
	hashAdapter hash.Adapter,
) FailureBuilder {
	out := failureBuilder{
		hashAdapter:     hashAdapter,
		pIndex:          nil,
		pCode:           nil,
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

// WithIndex adds an index to the builder
func (app *failureBuilder) WithIndex(index uint) FailureBuilder {
	app.pIndex = &index
	return app
}

// WithCode adds a code to the builder
func (app *failureBuilder) WithCode(code uint) FailureBuilder {
	app.pCode = &code
	return app
}

// IsRaisedInLayer flags the builder as isRaisedInLayer
func (app *failureBuilder) IsRaisedInLayer() FailureBuilder {
	app.isRaisedInLayer = true
	return app
}

// Now builds a new Failure instance
func (app *failureBuilder) Now() (Failure, error) {
	if app.pCode == nil {
		return nil, errors.New("the code is mandatory in order to build a Failure instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the code is mandatory in order to build a Failure instance")
	}

	isRaisedInLayer := "false"
	if app.isRaisedInLayer {
		isRaisedInLayer = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(*app.pIndex))),
		[]byte(strconv.Itoa(int(*app.pCode))),
		[]byte(isRaisedInLayer),
	})

	if err != nil {
		return nil, err
	}

	return createFailure(*pHash, *app.pIndex, *app.pCode, app.isRaisedInLayer), nil
}
