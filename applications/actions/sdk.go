package actions

import (
	"github.com/steve-care-software/syntax/domain/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/bytes/trees"
	"github.com/steve-care-software/syntax/domain/commands"
	"github.com/steve-care-software/syntax/domain/compilers"
	"github.com/steve-care-software/syntax/domain/outputs"
	"github.com/steve-care-software/syntax/domain/programs"
	"github.com/steve-care-software/syntax/domain/programs/instructions/applications/modules"
)

// Builder represents the builder application
type Builder interface {
	Create() Builder
	WithModules(modules modules.Modules) Builder
	Now() (Application, error)
}

// Application represents an action application
type Application interface {
	Tokenize(grammar grammars.Grammar, values []byte) (trees.Tree, []byte, error)
	Extract(criteria criterias.Criteria, tree trees.Tree) ([]byte, error)
	Combine(trees []trees.Tree, includeChannels bool) ([]byte, error)
	Compile(compiler compilers.Compiler, script []byte) (commands.Commands, []byte, error)
	Program(commands commands.Commands) (programs.Program, error)
	Execute(input map[string]interface{}, program programs.Program) (outputs.Output, error)
}
