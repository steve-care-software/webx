package trees

import (
	"errors"
)

type valueBuilder struct {
	pContent *byte
	prefix   Trees
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		pContent: nil,
		prefix:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithContent adds a content to the builder
func (app *valueBuilder) WithContent(content byte) ValueBuilder {
	app.pContent = &content
	return app
}

// WithPrefix adds a prefix to the builder
func (app *valueBuilder) WithPrefix(prefix Trees) ValueBuilder {
	app.prefix = prefix
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.pContent == nil {
		return nil, errors.New("the content is mandatory in order to build a Value instance")
	}

	if app.prefix != nil {
		return createValueWithPrefix(*app.pContent, app.prefix), nil
	}

	return createValue(*app.pContent), nil
}
