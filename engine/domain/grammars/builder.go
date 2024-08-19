package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
)

type builder struct {
	pVersion  *uint
	root      elements.Element
	rules     rules.Rules
	blocks    blocks.Blocks
	omissions elements.Elements
}

func createBuilder() Builder {
	out := builder{
		pVersion:  nil,
		root:      nil,
		rules:     nil,
		blocks:    nil,
		omissions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version uint) Builder {
	app.pVersion = &version
	return app
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root elements.Element) Builder {
	app.root = root
	return app
}

// WithRules adds rules to the builder
func (app *builder) WithRules(rules rules.Rules) Builder {
	app.rules = rules
	return app
}

// WithBlocks adds a blocks to the builder
func (app *builder) WithBlocks(blocks blocks.Blocks) Builder {
	app.blocks = blocks
	return app
}

// WithOmissions add omissions to the builder
func (app *builder) WithOmissions(omissions elements.Elements) Builder {
	app.omissions = omissions
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.pVersion == nil {
		return nil, errors.New("the version is mandatory in order to build a Grammar instance")
	}

	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build a Grammar instance")
	}

	if app.rules == nil {
		return nil, errors.New("the rules is mandatory in order to build a Grammar instance")
	}

	if app.blocks == nil {
		return nil, errors.New("the blocks is mandatory in order to build a Grammar instance")
	}

	if app.omissions != nil {
		return createGrammarWithOmissions(*app.pVersion, app.root, app.rules, app.blocks, app.omissions), nil
	}

	return createGrammar(*app.pVersion, app.root, app.rules, app.blocks), nil
}
