package indexes

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
)

type indexBuilder struct {
	hashAdapter hash.Adapter
	name        string
	criteria    criterias.Criteria
}

func createIndexBuilder(
	hashAdapter hash.Adapter,
) IndexBuilder {
	out := indexBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		criteria:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *indexBuilder) Create() IndexBuilder {
	return createIndexBuilder(app.hashAdapter)
}

// WithName adds a name to the builder
func (app *indexBuilder) WithName(name string) IndexBuilder {
	app.name = name
	return app
}

// WithCriteria adds a criteria to the builder
func (app *indexBuilder) WithCriteria(criteria criterias.Criteria) IndexBuilder {
	app.criteria = criteria
	return app
}

// Now builds a new Index instance
func (app *indexBuilder) Now() (Index, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Index instance")
	}

	if app.criteria == nil {
		return nil, errors.New("the criteria is mandatory in order to build an Index instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.criteria.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createIndex(*pHash, app.name, app.criteria), nil
}
