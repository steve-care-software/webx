package layers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type instructionsBuilder struct {
	hashAdapter hash.Adapter
	list        []Instruction
}

func createInstructionsBuilder(
	hashAdapter hash.Adapter,
) InstructionsBuilder {
	out := instructionsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionsBuilder) Create() InstructionsBuilder {
	return createInstructionsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *instructionsBuilder) WithList(list []Instruction) InstructionsBuilder {
	app.list = list
	return app
}

// Now builds a new Instructions instance
func (app *instructionsBuilder) Now() (Instructions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	data := [][]byte{}
	for _, oneIns := range app.list {
		data = append(data, oneIns.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createInstructions(*pHash, app.list), nil
}
