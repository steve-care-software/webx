package trees

import "errors"

type blockBuilder struct {
	lines []Line
}

func createBlockBuilder() BlockBuilder {
	out := blockBuilder{
		lines: nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder()
}

// WithLines add lines to the builder
func (app *blockBuilder) WithLines(lines []Line) BlockBuilder {
	app.lines = lines
	return app
}

// Now builds a new Block instance
func (app *blockBuilder) Now() (Block, error) {
	if app.lines != nil && len(app.lines) <= 0 {
		app.lines = nil
	}

	if app.lines == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Block instance")
	}

	var successful Line
	for _, oneLine := range app.lines {
		if oneLine.IsSuccessful() {
			successful = oneLine
			break
		}
	}

	if successful != nil {
		return createBlockWithSuccessful(app.lines, successful), nil
	}

	return createBlock(app.lines), nil
}
