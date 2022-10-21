package compilers

import (
	"github.com/steve-care-software/webx/applications/creates"
	"github.com/steve-care-software/webx/applications/criterias"
	"github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/applications/interpreters"
	"github.com/steve-care-software/webx/domain/compilers"
	"github.com/steve-care-software/webx/domain/compilers/outputs"
)

// NewApplication creates a new application
func NewApplication(
	createApp creates.Application,
) Application {
	grammarApp := grammars.NewApplication()
	criteriaApp := criterias.NewApplication()
	interpreterApp := interpreters.NewApplication()
	outputBuilder := outputs.NewBuilder()
	return createApplication(
		grammarApp,
		criteriaApp,
		interpreterApp,
		createApp,
		outputBuilder,
	)
}

// Application represents the compiler application
type Application interface {
	Execute(compiler compilers.Compiler, script []byte) (outputs.Output, error)
}

/*
	The output should contains instructions + remaining (optional)
*/
