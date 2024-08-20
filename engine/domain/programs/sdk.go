package programs

import (
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
)

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithRoot(root elements.Element) Builder
	WithInstructions(instructions instructions.Instructions) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Grammar() grammars.Grammar
	Root() elements.Element
	Instructions() instructions.Instructions
}
