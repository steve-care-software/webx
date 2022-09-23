package grammars

import "errors"

type suiteBuilder struct {
	valid   []byte
	invalid []byte
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		valid:   nil,
		invalid: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder()
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
	if app.valid != nil {
		return createSuiteWithValid(app.valid), nil
	}

	if app.invalid != nil {
		return createSuiteWithInvalid(app.invalid), nil
	}

	return nil, errors.New("the Suite is invalid")

}
