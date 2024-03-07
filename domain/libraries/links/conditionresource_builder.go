package links

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type conditionResourceBuilder struct {
	hashAdapter    hash.Adapter
	code           uint
	isRaiseInLayer bool
}

func createConditionResourceBuilder(
	hashAdapter hash.Adapter,
) ConditionResourceBuilder {
	out := conditionResourceBuilder{
		hashAdapter:    hashAdapter,
		code:           0,
		isRaiseInLayer: false,
	}

	return &out
}

// Create initializes the builder
func (app *conditionResourceBuilder) Create() ConditionResourceBuilder {
	return createConditionResourceBuilder(
		app.hashAdapter,
	)
}

// WithCode adds a code to the builder
func (app *conditionResourceBuilder) WithCode(code uint) ConditionResourceBuilder {
	app.code = code
	return app
}

// IsRaisedInLayer flags the builder as isRaisedInLayer
func (app *conditionResourceBuilder) IsRaisedInLayer() ConditionResourceBuilder {
	app.isRaiseInLayer = true
	return app
}

// Now builds a new ConditionResource instance
func (app *conditionResourceBuilder) Now() (ConditionResource, error) {
	if app.code <= 0 {
		return nil, errors.New("the code is mandatory in order to build a ConditionResource instance")
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

	return createConditionResource(*pHash, app.code, app.isRaiseInLayer), nil
}
