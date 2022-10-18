package compilers

import (
	"github.com/steve-care-software/syntax/applications/engines/creates"
	"github.com/steve-care-software/syntax/applications/engines/criterias"
	"github.com/steve-care-software/syntax/applications/engines/grammars"
	program_application "github.com/steve-care-software/syntax/applications/engines/programs"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/compilers/outputs"
)

// NewApplication creates a new application
func NewApplication(
	createApp creates.Application,
) Application {
	grammarApp := grammars.NewApplication()
	criteriaApp := criterias.NewApplication()

	modules, err := createApp.Modules().Execute()
	if err != nil {
		panic(err)
	}

	programApp, err := program_application.NewBuilder().Create().WithModules(modules).Now()
	if err != nil {
		panic(err)
	}

	outputBuilder := outputs.NewBuilder()
	return createApplication(
		grammarApp,
		criteriaApp,
		programApp,
		createApp,
		outputBuilder,
	)
}

// Application represents the compiler application
type Application interface {
	Execute(compiler compilers.Compiler, script []byte) (outputs.Output, error)
}
