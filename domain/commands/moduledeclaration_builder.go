package commands

import (
	"errors"

	"github.com/steve-care-software/webx/domain/criterias"
)

type moduleDeclarationBuilder struct {
	name criterias.Criteria
}

func createModuleDeclarationBuilder() ModuleDeclarationBuilder {
	out := moduleDeclarationBuilder{
		name: nil,
	}

	return &out
}

// Create initializes the builder
func (app *moduleDeclarationBuilder) Create() ModuleDeclarationBuilder {
	return createModuleDeclarationBuilder()
}

// WithName adds a name to the builder
func (app *moduleDeclarationBuilder) WithName(name criterias.Criteria) ModuleDeclarationBuilder {
	app.name = name
	return app
}

// Now builds a new ModuleDeclaration instance
func (app *moduleDeclarationBuilder) Now() (ModuleDeclaration, error) {
	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build a ModuleDeclaration instance")
	}

	return createModuleDeclaration(app.name), nil
}
