package transpilers

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/replacements"
)

// Application represents the transpiler application
type Application interface {
	Lex(input []byte) ([]byte, error)
	Parse(lexedInput []byte) (replacements.Replacements, error)
	Compile(replacements replacements.Replacements) (asts.AST, error)
}
