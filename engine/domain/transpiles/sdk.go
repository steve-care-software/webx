package transpiles

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/transpiles/blocks"
)

// Transpile represents a transpile
type Transpile interface {
	Origin() grammars.Grammar
	Target() grammars.Grammar
	Blocks() blocks.Blocks
	Root() string
}
