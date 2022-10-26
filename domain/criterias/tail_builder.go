package criterias

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type tailBuilder struct {
	hashAdapter hash.Adapter
	name        string
	delimiter   Delimiter
}

func createTailBuilder(
	hashAdapter hash.Adapter,
) TailBuilder {
	out := tailBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		delimiter:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *tailBuilder) Create() TailBuilder {
	return createTailBuilder(app.hashAdapter)
}

// WithName adds a name to the builder
func (app *tailBuilder) WithName(name string) TailBuilder {
	app.name = name
	return app
}

// WithDelimiter adds a delimiter to the builder
func (app *tailBuilder) WithDelimiter(delimiter Delimiter) TailBuilder {
	app.delimiter = delimiter
	return app
}

// Now builds a new Tail instance
func (app *tailBuilder) Now() (Tail, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Tail instance")
	}

	data := [][]byte{
		[]byte(app.name),
	}

	if app.delimiter != nil {
		data = append(data, app.delimiter.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.delimiter != nil {
		return createTailWithDelimiter(*pHash, app.name, app.delimiter), nil
	}

	return createTail(*pHash, app.name), nil
}
