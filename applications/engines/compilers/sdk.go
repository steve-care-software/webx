package compilers

import (
	"github.com/steve-care-software/syntax/applications/engines/creates"
	"github.com/steve-care-software/syntax/applications/engines/criterias"
	"github.com/steve-care-software/syntax/applications/engines/grammars"
	"github.com/steve-care-software/syntax/applications/engines/interpreters"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/compilers/outputs"
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
