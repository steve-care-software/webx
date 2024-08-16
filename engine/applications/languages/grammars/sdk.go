package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/stacks"
)

// Application represents the grammar application
type Application interface {
	Lex(input []byte) ([]byte, error)
	Parse(lexedInput []byte) (grammars.Grammar, error)
	Compile(grammar grammars.Grammar) ([]byte, error)
	Interpret(byteCode []byte) (stacks.Stack, error)
}
