package elements

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/logics"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	logic       logics.Logic
	condition   conditions.Condition
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		logic:       nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithLogic adds a logic to the builder
func (app *elementBuilder) WithLogic(logic logics.Logic) ElementBuilder {
	app.logic = logic
	return app
}

// WithCondition adds a condition to the builder
func (app *elementBuilder) WithCondition(condition conditions.Condition) ElementBuilder {
	app.condition = condition
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.logic == nil {
		return nil, errors.New("the logic is mandatory in order to build an Element instance")
	}

	data := [][]byte{
		app.logic.Hash().Bytes(),
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.condition != nil {
		return createElementWithCondition(*pHash, app.logic, app.condition), nil
	}

	return createElement(*pHash, app.logic), nil
}
