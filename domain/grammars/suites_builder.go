package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type suitesBuilder struct {
	hashAdapter hash.Adapter
	list        []Suite
}

func createSuitesBuilder(
	hashAdapter hash.Adapter,
) SuitesBuilder {
	out := suitesBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *suitesBuilder) Create() SuitesBuilder {
	return createSuitesBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *suitesBuilder) WithList(list []Suite) SuitesBuilder {
	app.list = list
	return app
}

// Now builds a new Suites instance
func (app *suitesBuilder) Now() (Suites, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Suite in order to build a Suites instance")
	}

	data := [][]byte{}
	for _, oneSuite := range app.list {
		data = append(data, oneSuite.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createSuites(*pHash, app.list), nil
}
