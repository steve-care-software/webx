package tokens

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type lineBuilder struct {
	elements []hash.Hash
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		elements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithElements add elements to the builder
func (app *lineBuilder) WithElements(elements []hash.Hash) LineBuilder {
	app.elements = elements
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.elements != nil && len(app.elements) <= 0 {
		app.elements = nil
	}

	if app.elements == nil {
		return nil, errors.New("there must be at least 1 Element in order to build a Line instance")
	}

	return createLine(app.elements), nil
}
