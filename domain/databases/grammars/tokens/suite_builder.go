package tokens

import "errors"

type suiteBuilder struct {
	isValid bool
	content []byte
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		isValid: false,
		content: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder()
}

// WithContent adds content to the builder
func (app *suiteBuilder) WithContent(content []byte) SuiteBuilder {
	app.content = content
	return app
}

// IsValid flags the builder as valid
func (app *suiteBuilder) IsValid() SuiteBuilder {
	app.isValid = true
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.content != nil && len(app.content) <= 0 {
		app.content = nil
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Suite instance")
	}

	return createSuite(app.isValid, app.content), nil
}
