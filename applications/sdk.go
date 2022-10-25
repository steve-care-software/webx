package applications

import (
	"github.com/steve-care-software/webx/applications/creates"
	"github.com/steve-care-software/webx/applications/criterias"
	"github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/applications/instructions"
	"github.com/steve-care-software/webx/applications/interpreters"
	"github.com/steve-care-software/webx/applications/programs"
	"github.com/steve-care-software/webx/domain/compilers"
)

// NewApplication creates a new application
func NewApplication(createApp creates.Application) Application {
	criteriaApp := criterias.NewApplication()
	grammarApp := grammars.NewApplication()
	interpreterApp := interpreters.NewApplication()
	instructionApp := instructions.NewApplication()

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
		instructionApp,
	)
}

// Application represents an engine application
type Application interface {
	Create() creates.Application
	Criteria() criterias.Application
	Grammar() grammars.Application
	Interpreter() interpreters.Application
	Instruction() instructions.Application
	Program() programs.Application
	ParseThenInterpret(input map[string]interface{}, script []byte) (map[string]interface{}, []byte, error)
	CompileThenParseThenInterpret(input map[string]interface{}, compiler compilers.Compiler, script []byte) (map[string]interface{}, []byte, []byte, error)
}
