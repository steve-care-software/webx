package programs

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

// Application represents the program application
type Application interface {
	Lex(grammar grammars.Grammar, input []byte) ([]byte, error)
	Parse(grammar grammars.Grammar, input []byte) (asts.AST, error)
	Compile(ast asts.AST) ([]byte, error)
	Interpret(bytecode []byte) (stacks.Stack, error)
}
