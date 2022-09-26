package trees

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
)

type lineBuilder struct {
	grammar  grammars.Line
	elements Elements
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		grammar:  nil,
		elements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *lineBuilder) WithGrammar(grammar grammars.Line) LineBuilder {
	app.grammar = grammar
	return app
}

// WithElements add elements to the builder
func (app *lineBuilder) WithElements(elements Elements) LineBuilder {
	app.elements = elements
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Line instance")
	}

	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Line instance")
	}

	return createLine(app.grammar, app.elements), nil
}
