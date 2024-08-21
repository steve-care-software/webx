package programs

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/syscalls"
)

type program struct {
	grammar      grammars.Grammar
	root         elements.Element
	instructions instructions.Instructions
	syscalls     syscalls.Syscalls
}

func createProgram(
	grammar grammars.Grammar,
	root elements.Element,
	instructions instructions.Instructions,
) Program {
	return createProgramInternally(grammar, root, instructions, nil)
}

func createProgramWithSyscalls(
	grammar grammars.Grammar,
	root elements.Element,
	instructions instructions.Instructions,
	syscalls syscalls.Syscalls,
) Program {
	return createProgramInternally(grammar, root, instructions, syscalls)
}

func createProgramInternally(
	grammar grammars.Grammar,
	root elements.Element,
	instructions instructions.Instructions,
	syscalls syscalls.Syscalls,
) Program {
	out := program{
		grammar:      grammar,
		root:         root,
		instructions: instructions,
		syscalls:     syscalls,
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

// HasSyscalls returns true if there is syscalls, false otherwise
func (obj *program) HasSyscalls() bool {
	return obj.syscalls != nil
}

// Syscalls returns the syscalls, if any
func (obj *program) Syscalls() syscalls.Syscalls {
	return obj.syscalls
}
