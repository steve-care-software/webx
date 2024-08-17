package programs

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

// Application represents the program application
type Application interface {
	Parse(grammar grammars.Grammar, input []byte) (asts.AST, error)
	Compile(ast asts.AST) ([]byte, error)
	Decompile(byteCode []byte) (asts.AST, error)
	Compose(grammar grammars.Grammar, ast asts.AST) ([]byte, error)
	Interpret(bytecode []byte) (stacks.Stack, error)
}
