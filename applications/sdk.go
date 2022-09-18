package applications

import (
	"github.com/steve-care-software/logics/domain/bytes/criterias"
	"github.com/steve-care-software/logics/domain/bytes/grammars"
	"github.com/steve-care-software/logics/domain/bytes/trees"
	"github.com/steve-care-software/logics/domain/compilers"
	"github.com/steve-care-software/logics/domain/outputs"
	"github.com/steve-care-software/logics/domain/programs"
)

// Application represents the syntax application
type Application interface {
	Tokenize(grammar grammars.Grammar, values []byte) (trees.Tree, error)
	Extract(criteria criterias.Criteria, tree trees.Tree) ([]byte, error)
	Combine(trees []trees.Tree) ([]byte, error)
	Compile(compiler compilers.Compiler, script string) (programs.Program, error)
	Execute(params map[string]interface{}, program programs.Program) (outputs.Output, error)
}
