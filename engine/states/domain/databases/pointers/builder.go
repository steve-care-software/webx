package pointers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/states/domain/databases/metadatas"
)

type builder struct {
	hashAdapter hash.Adapter
	head        hash.Hash
	metaData    metadatas.MetaData
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		head:        nil,
		metaData:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head hash.Hash) Builder {
	app.head = head
	return app
}

// WithMetaData adds metaData to the builder
func (app *builder) WithMetaData(metaData metadatas.MetaData) Builder {
	app.metaData = metaData
	return app
}

// Now builds a new Pointer instance
func (app *builder) Now() (Pointer, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Pointer instance")
	}

	if app.metaData == nil {
		return nil, errors.New("the metaData is mandatory in order to build a Pointer instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.head.Bytes(),
		app.metaData.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createPointer(*pHash, app.head, app.metaData), nil
}
