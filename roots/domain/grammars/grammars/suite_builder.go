package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	valid       []byte
	invalid     []byte
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		valid:       nil,
		invalid:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder(
		app.hashAdapter,
	)
}

// WithValid add valid bytes to the builder
func (app *suiteBuilder) WithValid(valid []byte) SuiteBuilder {
	app.valid = valid
	return app
}

// WithInvalid add invalid bytes to the builder
func (app *suiteBuilder) WithInvalid(invalid []byte) SuiteBuilder {
	app.invalid = invalid
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {

	isValidStr := "false"
	if app.valid != nil {
		isValidStr = "true"
	}

	data := [][]byte{
		[]byte(isValidStr),
	}

	if app.valid != nil {
		data = append(data, app.valid)
	}

	if app.invalid != nil {
		data = append(data, app.invalid)
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.valid != nil {
		return createSuiteWithValid(*pHash, app.valid), nil
	}

	if app.invalid != nil {
		return createSuiteWithInvalid(*pHash, app.invalid), nil
	}

	return nil, errors.New("the Suite is invalid")

}
