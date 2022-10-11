package engines

import (
	"github.com/steve-care-software/syntax/applications/engines/creates"
	"github.com/steve-care-software/syntax/applications/engines/criterias"
	"github.com/steve-care-software/syntax/applications/engines/grammars"
	"github.com/steve-care-software/syntax/applications/engines/interpreters"
	"github.com/steve-care-software/syntax/applications/engines/programs"
)

// NewApplication creates a new application
func NewApplication(createApp creates.Application) Application {
	criteriaApp := criterias.NewApplication()
	grammarApp := grammars.NewApplication()
	interpreterApp := interpreters.NewApplication()

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
		criteriaApp,
		grammarApp,
		interpreterApp,
		programApp,
	)
}

// Application represents an engine application
type Application interface {
	Create() creates.Application
	Criteria() criterias.Application
	Grammar() grammars.Application
	Interpreter() interpreters.Application
	Program() programs.Application
	ParseThenInterpret(input map[string]interface{}, script []byte) (map[string]interface{}, []byte, error)
}
