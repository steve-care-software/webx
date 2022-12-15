package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type lineBuilder struct {
	hashAdapter hash.Adapter
	elements    []Element
}

func createLineBuilder(
	hashAdapter hash.Adapter,
) LineBuilder {
	out := lineBuilder{
		hashAdapter: hashAdapter,
		elements:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder(
		app.hashAdapter,
	)
}

// WithElements add elements to the builder
func (app *lineBuilder) WithElements(elements []Element) LineBuilder {
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

	data := [][]byte{}
	for _, oneElement := range app.elements {
		data = append(data, oneElement.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createLine(*pHash, app.elements), nil
}
