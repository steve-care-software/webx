package programs

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
)

type builder struct {
	grammar grammars.Grammar
	root    instructions.Element
}

func createBuilder() Builder {
	out := builder{
		grammar: nil,
		root:    nil,
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

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root instructions.Element) Builder {
	app.root = root
	return app
}

// Now builds a new Program
func (app *builder) Now() (Program, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Program instance")
	}

	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build a Program instance")
	}

	return createProgram(
		app.grammar,
		app.root,
	), nil
}
