package elements

import (
	"errors"

	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	name        string
	criteria    criterias.Criteria
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		criteria:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(app.hashAdapter)
}

// WithName adds a name to the builder
func (app *elementBuilder) WithName(name string) ElementBuilder {
	app.name = name
	return app
}

// WithCriteria adds a criteria to the builder
func (app *elementBuilder) WithCriteria(criteria criterias.Criteria) ElementBuilder {
	app.criteria = criteria
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Element instance")
	}

	if app.criteria == nil {
		return nil, errors.New("the criteria is mandatory in order to build an Element instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.criteria.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createElement(*pHash, app.name, app.criteria), nil
}
