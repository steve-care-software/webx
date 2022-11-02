package applications

import (
	"github.com/steve-care-software/webx/applications/creates"
	"github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/applications/interpreters"
	"github.com/steve-care-software/webx/applications/programs"
	selector_application "github.com/steve-care-software/webx/applications/selectors"
	"github.com/steve-care-software/webx/domain/compilers"
)

// NewApplication creates a new application
func NewApplication(createApp creates.Application) Application {
	grammarApp := grammars.NewApplication()
	interpreterApp := interpreters.NewApplication()
	selectorApp := selector_application.NewApplication()
	modules, err := createApp.Modules().Execute()
	if err != nil {
		panic(err)
	}

	programApp, err := programs.NewBuilder().Create().WithModules(modules).Now()
	if err != nil {
		panic(err)
	}

	return createApplication(
		createApp,
		grammarApp,
		interpreterApp,
		programApp,
		selectorApp,
	)
}

// Application represents an engine application
type Application interface {
	Create() creates.Application
	Grammar() grammars.Application
	Interpreter() interpreters.Application
	Program() programs.Application
	Selector() selector_application.Application
	ParseThenInterpret(input map[string]interface{}, script []byte) (map[string]interface{}, []byte, error)
	CompileThenParseThenInterpret(input map[string]interface{}, compiler compilers.Compiler, script []byte) (map[string]interface{}, []byte, []byte, error)
}
