package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type blockBuilder struct {
	hashAdapter hash.Adapter
	lines       []Line
}

func createBlockBuilder(
	hashAdapter hash.Adapter,
) BlockBuilder {
	out := blockBuilder{
		hashAdapter: hashAdapter,
		lines:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{}
	for _, oneLine := range app.lines {
		data = append(data, oneLine.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createBlock(*pHash, app.lines), nil
}
