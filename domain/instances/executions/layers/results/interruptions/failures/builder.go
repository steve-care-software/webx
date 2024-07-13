package failures

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter     hash.Adapter
	pIndex          *uint
	pCode           *uint
	isRaisedInLayer bool
	message         string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:     hashAdapter,
		pIndex:          nil,
		pCode:           nil,
		isRaisedInLayer: false,
		message:         "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithCode adds a code to the builder
func (app *builder) WithCode(code uint) Builder {
	app.pCode = &code
	return app
}

// IsRaisedInLayer flags the builder as isRaisedInLayer
func (app *builder) IsRaisedInLayer() Builder {
	app.isRaisedInLayer = true
	return app
}

// WithMessage adds a message to the builder
func (app *builder) WithMessage(message string) Builder {
	app.message = message
	return app
}

// Now builds a new Failure instance
func (app *builder) Now() (Failure, error) {
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

	data := [][]byte{
		[]byte(strconv.Itoa(int(*app.pIndex))),
		[]byte(strconv.Itoa(int(*app.pCode))),
		[]byte(isRaisedInLayer),
	}

	if app.message != "" {
		data = append(data, []byte(app.message))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.message != "" {
		return createFailureWithMessage(*pHash, *app.pIndex, *app.pCode, app.isRaisedInLayer, app.message), nil
	}

	return createFailure(*pHash, *app.pIndex, *app.pCode, app.isRaisedInLayer), nil
}
