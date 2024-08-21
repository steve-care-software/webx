package programs

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
)

type program struct {
	grammar      grammars.Grammar
	root         elements.Element
	instructions instructions.Instructions
}

func createProgram(
	grammar grammars.Grammar,
	root elements.Element,
	instructions instructions.Instructions,
) Program {
	out := program{
		grammar:      grammar,
		root:         root,
		instructions: instructions,
	}

	return &out
}

// Grammar returns the grammar
func (obj *program) Grammar() grammars.Grammar {
	return obj.grammar
}

// Root returns the root
func (obj *program) Root() elements.Element {
	return obj.root
}

// Instructions returns the instructions
func (obj *program) Instructions() instructions.Instructions {
	return obj.instructions
}
