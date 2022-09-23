package trees

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/bytes/grammars"
)

type builder struct {
	grammar grammars.Token
	block   Block
}

func createBuilder() Builder {
	out := builder{
		grammar: nil,
		block:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar grammars.Token) Builder {
	app.grammar = grammar
	return app
}

// WithBlock adds a block to the builder
func (app *builder) WithBlock(block Block) Builder {
	app.block = block
	return app
}

// Now builds a new Tree instance
func (app *builder) Now() (Tree, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Tree instance")
	}

	if app.block == nil {
		return nil, errors.New("the block is mandatory in order to build a Tree instance")
	}

	return createTree(app.grammar, app.block), nil
}
