package selectors

import (
	"errors"

	"github.com/steve-care-software/webx/grammars/domain/grammars"
)

type builder struct {
	grammar grammars.Grammar
	token   Token
	inside  Inside
	fn      SelectorFn
}

func createBuilder() Builder {
	out := builder{
		grammar: nil,
		token:   nil,
		inside:  nil,
		fn:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar grammars.Grammar) Builder {
	app.grammar = grammar
	return app
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token Token) Builder {
	app.token = token
	return app
}

// WithInside adds an inside to the builder
func (app *builder) WithInside(inside Inside) Builder {
	app.inside = inside
	return app
}

// WithFn adds a func to the builder
func (app *builder) WithFn(fn SelectorFn) Builder {
	app.fn = fn
	return app
}

// Now builds a new Selector instance
func (app *builder) Now() (Selector, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Selector instance")
	}

	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Selector instance")
	}

	if app.inside == nil {
		return nil, errors.New("the inside is mandatory in order to build a Selector instance")
	}

	if app.fn == nil {
		return nil, errors.New("the func is mandatory in order to build a Selector instance")
	}

	return createSelector(app.grammar, app.token, app.inside, app.fn), nil
}
