package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type comparisonsBuilder struct {
	hashAdapter hash.Adapter
	list        []Comparison
}

func createComparisonsBuilder(
	hashAdapter hash.Adapter,
) ComparisonsBuilder {
	out := comparisonsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *comparisonsBuilder) Create() ComparisonsBuilder {
	return createComparisonsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *comparisonsBuilder) WithList(list []Comparison) ComparisonsBuilder {
	app.list = list
	return app
}

// Now builds a new Comparisons instance
func (app *comparisonsBuilder) Now() (Comparisons, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Comparison in order to build a Comparisons instance")
	}

	data := [][]byte{}
	for _, oneComparison := range app.list {
		data = append(data, oneComparison.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createComparisons(*pHash, app.list), nil
}
