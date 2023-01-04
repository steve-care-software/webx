package applications

import (
	grammars_application "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/interpreters/domain/results"
	programs_application "github.com/steve-care-software/webx/programs/applications"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
	selectors_application "github.com/steve-care-software/webx/selectors/applications"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// FetchModulesFn returns the modules
type FetchModulesFn func() (modules.Modules, error)

// NewBuilder creates a new builder instance
func NewBuilder(
	nameBytesToStringFn programs_application.NameBytesToString,
) Builder {
	resultBuilder := results.NewBuilder()
	grammarApp := grammars_application.NewApplication()
	selectorApp := selectors_application.NewApplication()
	programApp := programs_application.NewApplication(
		nameBytesToStringFn,
	)

	return createBuilder(
		resultBuilder,
		grammarApp,
		selectorApp,
		programApp,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithSelector(selector selectors.Selector) Builder
	WithModulesFn(modulesFn FetchModulesFn) Builder
	Now() (Application, error)
}

// Application represents an interpreter application
type Application interface {
	ParseThenInterpret(input []interface{}, script []byte) (results.Result, error)
}
