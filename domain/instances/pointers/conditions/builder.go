package conditions

import (
	"errors"

	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	resource    Resource
	comparisons Comparisons
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		resource:    nil,
		comparisons: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithResource adds a resource to the builder
func (app *builder) WithResource(resource Resource) Builder {
	app.resource = resource
	return app
}

// WithComparisons adds a comparisons to the builder
func (app *builder) WithComparisons(comparisons Comparisons) Builder {
	app.comparisons = comparisons
	return app
}

// Now builds a new Condition instance
func (app *builder) Now() (Condition, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build a Condition instance")
	}

	data := [][]byte{
		app.resource.Hash().Bytes(),
	}

	if app.comparisons != nil {
		data = append(data, app.comparisons.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.comparisons != nil {
		return createConditionWithComparisons(*pHash, app.resource, app.comparisons), nil
	}

	return createCondition(*pHash, app.resource), nil
}
