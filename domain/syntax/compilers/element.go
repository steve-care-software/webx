package compilers

import (
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
)

type element struct {
	grammar    grammars.Grammar
	program    programs.Program
	parameters Parameters
}

func createElement(
	grammar grammars.Grammar,
	program programs.Program,
	parameters Parameters,
) Element {
	out := element{
		grammar:    grammar,
		program:    program,
		parameters: parameters,
	}

	return &out
}

// Grammar returns the grammar
func (obj *element) Grammar() grammars.Grammar {
	return obj.grammar
}

// Program returns the program
func (obj *element) Program() programs.Program {
	return obj.program
}

// Parameters returns the parameters
func (obj *element) Parameters() Parameters {
	return obj.parameters
}
