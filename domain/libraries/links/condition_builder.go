package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions/resources"
)

type conditionBuilder struct {
	hashAdapter hash.Adapter
	resource    resources.Resource
	next        ConditionValue
}

func createConditionBuilder(
	hashAdapter hash.Adapter,
) ConditionBuilder {
	out := conditionBuilder{
		hashAdapter: hashAdapter,
		resource:    nil,
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder(
		app.hashAdapter,
	)
}

// WithResource adds a resource builder
func (app *conditionBuilder) WithResource(resource resources.Resource) ConditionBuilder {
	app.resource = resource
	return app
}

// WithNext adds a next value to the builder builder
func (app *conditionBuilder) WithNext(next ConditionValue) ConditionBuilder {
	app.next = next
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build an Condition instance")
	}

	data := [][]byte{
		app.resource.Hash().Bytes(),
	}
	if app.next != nil {
		data = append(data, app.next.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.next != nil {
		return createConditionWithNext(*pHash, app.resource, app.next), nil
	}

	return createCondition(*pHash, app.resource), nil
}
