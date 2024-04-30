package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/resources"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	condition   Condition
	resource    resources.Resource
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		condition:   nil,
		resource:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithCondition adds a condition to the builder
func (app *elementBuilder) WithCondition(condition Condition) ElementBuilder {
	app.condition = condition
	return app
}

// WithResource adds a resource to the builder
func (app *elementBuilder) WithResource(resource resources.Resource) ElementBuilder {
	app.resource = resource
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	data := [][]byte{}
	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	if app.resource != nil {
		data = append(data, app.resource.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Element is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.condition != nil {
		return createElementWithCondition(*pHash, app.condition), nil
	}

	return createElementWithResource(*pHash, app.resource), nil
}
