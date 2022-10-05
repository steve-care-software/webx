package compilers

import (
	coverage_application "github.com/steve-care-software/syntax/applications/grammars/coverages"
	"github.com/steve-care-software/syntax/applications/interpreters"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions/applications/modules"
)

// CompilerFn represents the compiler func
type CompilerFn func(tree trees.Tree) (compilers.Compiler, error)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithModules(modules modules.Modules) Builder
	Now() (Application, error)
}

// Application represents the compiler application
type Application interface {
	Coverage() coverage_application.Application
	Instance(input []byte) (compilers.Compiler, error)
	Execute(compiler compilers.Compiler, script []byte) (commands.Commands, []byte, error)
	Interpreter(commands commands.Commands) (interpreters.Application, error)
}
