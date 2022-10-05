package interpreters

import (
	creates_command "github.com/steve-care-software/syntax/applications/engines/creates/commands"
	creates_grammar "github.com/steve-care-software/syntax/applications/engines/creates/grammars"
	creates_module "github.com/steve-care-software/syntax/applications/engines/creates/modules"
	"github.com/steve-care-software/syntax/applications/engines/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/outputs"
)

// NewBuilder creates a new interpreter builder
func NewBuilder() Builder {
	grammarApp := grammars.NewApplication()
	return createBuilder(nil, grammarApp)
}

// Builder represents the interpreter builder
type Builder interface {
	Create() Builder
	WithScript(script []byte) Builder
	WithGrammar(grammar creates_grammar.Application) Builder
	WithCommand(command creates_command.Application) Builder
	WithModules(modules creates_module.Application) Builder
	Now() (Application, []byte, error)
}

// Application represents an interpreter application
type Application interface {
	Execute(input map[string]interface{}) (outputs.Output, error)
}
