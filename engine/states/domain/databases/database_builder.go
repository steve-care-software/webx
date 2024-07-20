package databases

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/domain/databases/metadatas"
	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	head        commits.Commit
	metaData    metadatas.MetaData
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	return &builder{
		hashAdapter: hashAdapter,
		head:        nil,
		metaData:    nil,
	}
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithHead adds a head commit to the builder
func (app *builder) WithHead(head commits.Commit) Builder {
	app.head = head
	return app
}

// WithMetaData adds metadata to the builder
func (app *builder) WithMetaData(metaData metadatas.MetaData) Builder {
	app.metaData = metaData
	return app
}

// Now builds a new Database instance
func (app *builder) Now() (Database, error) {
	if app.head == nil {
		return nil, errors.New("the head commit is mandatory in order to build a Database instance")
	}

	if app.metaData == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Database instance")
	}

	data := [][]byte{
		app.head.Hash().Bytes(),
		app.metaData.Hash().Bytes(),
	}

	dbHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createDatabase(*dbHash, app.head, app.metaData), nil
}
