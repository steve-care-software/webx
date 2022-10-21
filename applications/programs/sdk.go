package programs

import (
	"github.com/steve-care-software/webx/domain/commands"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/programs/applications/modules"
	"github.com/steve-care-software/webx/domain/programs/outputs"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithModules(modules modules.Modules) Builder
	Now() (Application, error)
}

// Application represents a program application
type Application interface {
	Execute(grammar grammars.Grammar, command commands.Command, script []byte) (outputs.Output, error)
}

/*
	Replace all inputs of the execute by instructions
	The output is only a program instance (no remaining, therefore no need for output)
*/
