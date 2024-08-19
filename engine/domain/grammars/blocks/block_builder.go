package blocks

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/suites"
)

type blockBuilder struct {
	name   string
	lines  lines.Lines
	suites suites.Suites
}

func createBlockBuilder() BlockBuilder {
	out := blockBuilder{
		name:   "",
		lines:  nil,
		suites: nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder()
}

// WithName adds a name to the builder
func (app *blockBuilder) WithName(name string) BlockBuilder {
	app.name = name
	return app
}

// WithLines add lines to the builder
func (app *blockBuilder) WithLines(lines lines.Lines) BlockBuilder {
	app.lines = lines
	return app
}

// WithSuites add suites to the builder
func (app *blockBuilder) WithSuites(suites suites.Suites) BlockBuilder {
	app.suites = suites
	return app
}

// Now builds a new Block instance
func (app *blockBuilder) Now() (Block, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Block instance")
	}

	if app.lines == nil {
		return nil, errors.New("the lines is mandatory in order to build a Block instance")
	}

	if app.suites != nil {
		return createBlockWithSuites(app.name, app.lines, app.suites), nil
	}

	return createBlock(app.name, app.lines), nil
}
