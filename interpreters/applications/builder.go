package applications

import (
	"errors"

	grammars_application "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/interpreters/domain/results"
	programs_application "github.com/steve-care-software/webx/programs/applications"
	selectors_application "github.com/steve-care-software/webx/selectors/applications"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

type builder struct {
	resultBuilder results.Builder
	grammarApp    grammars_application.Application
	selectorApp   selectors_application.Application
	programApp    programs_application.Application
	grammarIns    grammars.Grammar
	selectorIns   selectors.Selector
	modulesFn     FetchModulesFn
}

func createBuilder(
	resultBuilder results.Builder,
	grammarApp grammars_application.Application,
	selectorApp selectors_application.Application,
	programApp programs_application.Application,
) Builder {
	out := builder{
		resultBuilder: resultBuilder,
		grammarApp:    grammarApp,
		selectorApp:   selectorApp,
		programApp:    programApp,
		grammarIns:    nil,
		selectorIns:   nil,
		modulesFn:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.resultBuilder,
		app.grammarApp,
		app.selectorApp,
		app.programApp,
	)
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar grammars.Grammar) Builder {
	app.grammarIns = grammar
	return app
}

// WithSelector adds a selector to the builder
func (app *builder) WithSelector(selector selectors.Selector) Builder {
	app.selectorIns = selector
	return app
}

// WithModulesFn add modules func to the builder
func (app *builder) WithModulesFn(modulesFn FetchModulesFn) Builder {
	app.modulesFn = modulesFn
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.grammarIns == nil {
		return nil, errors.New("the grammar instance is mandatory in order to build an Application instance")
	}

	if app.selectorIns == nil {
		return nil, errors.New("the selector instance is mandatory in order to build an Application instance")
	}

	if app.modulesFn == nil {
		return nil, errors.New("the modules's Func is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.resultBuilder,
		app.grammarApp,
		app.selectorApp,
		app.programApp,
		app.grammarIns,
		app.selectorIns,
		app.modulesFn,
	), nil
}
