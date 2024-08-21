package programs

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
)

type builder struct {
	grammar      grammars.Grammar
	root         elements.Element
	instructions instructions.Instructions
}

func createBuilder() Builder {
	out := builder{
		grammar:      nil,
		root:         nil,
		instructions: nil,
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
func (app *builder) WithRoot(root elements.Element) Builder {
	app.root = root
	return app
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(instructions instructions.Instructions) Builder {
	app.instructions = instructions
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

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Program instance")
	}

	return createProgram(
		app.grammar,
		app.root,
		app.instructions,
	), nil
}
