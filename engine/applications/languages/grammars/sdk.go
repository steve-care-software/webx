package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars"
)

// Application represents the grammar application
type Application interface {
	Parse(lexedInput []byte) (grammars.Grammar, error)
	Compile(grammar grammars.Grammar) (asts.AST, error)
	Decompile(ast asts.AST) (grammars.Grammar, error)
	Compose(grammar grammars.Grammar) ([]byte, error)
}
