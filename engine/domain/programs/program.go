package programs

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
)

type program struct {
	grammar grammars.Grammar
	root    instructions.Element
}

func createProgram(
	grammar grammars.Grammar,
	root instructions.Element,
) Program {
	return createProgramInternally(grammar, root)
}

func createProgramInternally(
	grammar grammars.Grammar,
	root instructions.Element,
) Program {
	out := program{
		grammar: grammar,
		root:    root,
	}

	return &out
}

// Grammar returns the grammar
func (obj *program) Grammar() grammars.Grammar {
	return obj.grammar
}

// Root returns the root
func (obj *program) Root() instructions.Element {
	return obj.root
}
