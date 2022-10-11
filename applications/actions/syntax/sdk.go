package syntax

import (
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/trees"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/composers"
	"github.com/steve-care-software/syntax/domain/syntax/outputs"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions/applications/modules"
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
	Compose(composer composers.Composer) ([]byte, error)
	Program(commands commands.Commands) (programs.Program, error)
	Execute(input map[string]interface{}, program programs.Program) (outputs.Output, error)
}
