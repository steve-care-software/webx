package trees

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
)

type treeBuilder struct {
	grammar   grammars.Token
	block     Block
	suffix    Trees
	remaining []byte
}

func createTreeBuilder() TreeBuilder {
	out := treeBuilder{
		grammar:   nil,
		block:     nil,
		suffix:    nil,
		remaining: nil,
	}

	return &out
}

// Create initializes the treeBuilder
func (app *treeBuilder) Create() TreeBuilder {
	return createTreeBuilder()
}

// WithGrammar adds a grammar to the treeBuilder
func (app *treeBuilder) WithGrammar(grammar grammars.Token) TreeBuilder {
	app.grammar = grammar
	return app
}

// WithBlock adds a block to the treeBuilder
func (app *treeBuilder) WithBlock(block Block) TreeBuilder {
	app.block = block
	return app
}

// WithSuffix adds a suffix to the builder
func (app *treeBuilder) WithSuffix(suffix Trees) TreeBuilder {
	app.suffix = suffix
	return app
}

// WithRemaining adds a remaining to the builder
func (app *treeBuilder) WithRemaining(remaining []byte) TreeBuilder {
	app.remaining = remaining
	return app
}

// Now builds a new Tree instance
func (app *treeBuilder) Now() (Tree, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Tree instance")
	}

	if app.block == nil {
		return nil, errors.New("the block is mandatory in order to build a Tree instance")
	}

	if app.remaining != nil && len(app.remaining) <= 0 {
		app.remaining = nil
	}

	if app.remaining != nil && app.suffix != nil {
		return createTreeWithSuffixAndRemaining(app.grammar, app.block, app.suffix, app.remaining), nil
	}

	if app.remaining != nil {
		return createTreeWithRemaining(app.grammar, app.block, app.remaining), nil
	}

	if app.suffix != nil {
		return createTreeWithSuffix(app.grammar, app.block, app.suffix), nil
	}

	return createTree(app.grammar, app.block), nil
}
