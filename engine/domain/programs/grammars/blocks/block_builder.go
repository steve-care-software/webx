package blocks

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
)

type blockBuilder struct {
	name   string
	line   lines.Line
	lines  lines.Lines
	suites suites.Suites
}

func createBlockBuilder() BlockBuilder {
	out := blockBuilder{
		name:   "",
		line:   nil,
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

// WithLine add line to the builder
func (app *blockBuilder) WithLine(line lines.Line) BlockBuilder {
	app.line = line
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

	if app.line != nil && app.suites != nil {
		return createBlockWithLineAndSuites(app.name, app.line, app.suites), nil
	}

	if app.line != nil {
		return createBlockWithLine(app.name, app.line), nil
	}

	if app.lines != nil && app.suites != nil {
		return createBlockWithLinesAndSuites(app.name, app.lines, app.suites), nil
	}

	if app.lines != nil {
		return createBlockWithLines(app.name, app.lines), nil
	}

	return nil, errors.New("the Block is invalid")
}
