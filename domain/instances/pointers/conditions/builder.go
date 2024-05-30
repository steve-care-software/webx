package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
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

	if app.comparisons == nil {
		return nil, errors.New("the comparisons is mandatory in order to build a Condition instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.resource.Hash().Bytes(),
		app.comparisons.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createCondition(*pHash, app.resource, app.comparisons), nil
}
