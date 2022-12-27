package trees

import (
	"errors"

	"github.com/steve-care-software/webx/grammars/domain/grammars"
)

type lineBuilder struct {
	pIndex    *uint
	grammar   grammars.Line
	isReverse bool
	elements  Elements
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		pIndex:    nil,
		grammar:   nil,
		isReverse: false,
		elements:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithIndex adds an index to the builder
func (app *lineBuilder) WithIndex(index uint) LineBuilder {
	app.pIndex = &index
	return app
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

// IsReverse flags the builder as reverse
func (app *lineBuilder) IsReverse() LineBuilder {
	app.isReverse = true
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Line instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Line instance")
	}

	if app.elements != nil {
		return createLineWithElements(*app.pIndex, app.grammar, app.isReverse, app.elements), nil
	}

	return createLine(*app.pIndex, app.grammar, app.isReverse), nil
}
