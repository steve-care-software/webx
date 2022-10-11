package commands

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/criterias"
)

type applicationDeclarationBuilder struct {
	module criterias.Criteria
	name   criterias.Criteria
}

func createApplicationDeclarationBuilder() ApplicationDeclarationBuilder {
	out := applicationDeclarationBuilder{
		module: nil,
		name:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationDeclarationBuilder) Create() ApplicationDeclarationBuilder {
	return createApplicationDeclarationBuilder()
}

// WithModule adds a module to the builder
func (app *applicationDeclarationBuilder) WithModule(module criterias.Criteria) ApplicationDeclarationBuilder {
	app.module = module
	return app
}

// WithName adds a name to the builder
func (app *applicationDeclarationBuilder) WithName(name criterias.Criteria) ApplicationDeclarationBuilder {
	app.name = name
	return app
}

// Now builds a new ApplicationDeclaration instance
func (app *applicationDeclarationBuilder) Now() (ApplicationDeclaration, error) {
	if app.module == nil {
		return nil, errors.New("the module is mandatory in order to build an ApplicationDeclaration instance")
	}

	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build an ApplicationDeclaration instance")
	}

	return createApplicationDeclaration(app.module, app.name), nil
}
