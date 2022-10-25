package selections

import (
	"errors"

	"github.com/steve-care-software/webx/domain/trees"
)

type elementBuilder struct {
	value               trees.Element
	includeChannelBytes bool
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		value:               nil,
		includeChannelBytes: false,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithValue adds a value to the builder
func (app *elementBuilder) WithValue(value trees.Element) ElementBuilder {
	app.value = value
	return app
}

// IncludeChannelBytesflags the builder as includeChannelBytes
func (app *elementBuilder) IncludeChannelBytes() ElementBuilder {
	app.includeChannelBytes = true
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build an Element instance")
	}

	return createElement(app.value, app.includeChannelBytes), nil
}
