package transpilers

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/transpiles"
)

// Application represents the transpiler application
type Application interface {
	Parse(lexedInput []byte) (transpiles.Transpile, error)
	Compile(transpile transpiles.Transpile) (asts.AST, error)
	Decompile(ast asts.AST) (transpiles.Transpile, error)
	Compose(transpile transpiles.Transpile) ([]byte, error)
}
