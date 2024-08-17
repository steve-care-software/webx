package transpilers

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/replacements"
	"github.com/steve-care-software/webx/engine/domain/transpiles"
)

// Application represents the transpiler application
type Application interface {
	Lex(input []byte) ([]byte, error)
	Parse(lexedInput []byte) (transpiles.Transpile, error)
	Compile(replacements replacements.Replacements) (asts.AST, error)
}
