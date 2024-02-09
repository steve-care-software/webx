package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type conditionValueBuilder struct {
	hashAdapter hash.Adapter
	resource    ConditionResource
	condition   Condition
}

func createConditionValueBuilder(
	hashAdapter hash.Adapter,
) ConditionValueBuilder {
	out := conditionValueBuilder{
		hashAdapter: hashAdapter,
		resource:    nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionValueBuilder) Create() ConditionValueBuilder {
	return createConditionValueBuilder(
		app.hashAdapter,
	)
}

// WithResource adds a resource to the builder
func (app *conditionValueBuilder) WithResource(resource ConditionResource) ConditionValueBuilder {
	app.resource = resource
	return app
}

// WithCondition adds an condition to the builder
func (app *conditionValueBuilder) WithCondition(condition Condition) ConditionValueBuilder {
	app.condition = condition
	return app
}

// Now builds a new ConditionValue instance
func (app *conditionValueBuilder) Now() (ConditionValue, error) {
	data := [][]byte{}
	if app.resource != nil {
		data = append(data, app.resource.Hash().Bytes())
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the ConditionValue is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.resource != nil {
		return createConditionValueWithResource(*pHash, app.resource), nil
	}

	return createConditionValueWithCondition(*pHash, app.condition), nil
}
