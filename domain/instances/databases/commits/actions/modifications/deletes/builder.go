package deletes

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pIndex      *uint
	length      uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pIndex:      nil,
		length:      0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithLength adds a length to the builder
func (app *builder) WithLength(length uint) Builder {
	app.length = length
	return app
}

// Now builds a new Delete instance
func (app *builder) Now() (Delete, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Delete instance")
	}

	if app.length <= 0 {
		return nil, errors.New("the length is mandatory in order to build a Delete instance")

	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(*app.pIndex))),
		[]byte(strconv.Itoa(int(app.length))),
	})

	if err != nil {
		return nil, err
	}

	return createDelete(*pHash, *app.pIndex, app.length), nil
}
