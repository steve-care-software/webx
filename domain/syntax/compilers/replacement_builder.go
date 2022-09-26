package compilers

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
)

type replacementBuilder struct {
	name     []byte
	criteria criterias.Criteria
}

func createReplacementBuilder() ReplacementBuilder {
	out := replacementBuilder{
		name:     nil,
		criteria: nil,
	}

	return &out
}

// Create initializes the builder
func (app *replacementBuilder) Create() ReplacementBuilder {
	return createReplacementBuilder()
}

// WithName adds a name to the builder
func (app *replacementBuilder) WithName(name []byte) ReplacementBuilder {
	app.name = name
	return app
}

// WithCriteria adds a criteria to the builder
func (app *replacementBuilder) WithCriteria(criteria criterias.Criteria) ReplacementBuilder {
	app.criteria = criteria
	return app
}

// Now builds a new Replacement instance
func (app *replacementBuilder) Now() (Replacement, error) {
	if app.name != nil && len(app.name) <= 0 {
		app.name = nil
	}

	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build a Replacement instance")
	}

	if app.criteria == nil {
		return nil, errors.New("the criteria is mandatory in order to build a Replacement instance")
	}

	return createReplacement(app.name, app.criteria), nil
}
