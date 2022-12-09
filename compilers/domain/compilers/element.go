package compilers

import (
	"github.com/steve-care-software/webx/grammars/domain/grammars"
)

type element struct {
	grammar    grammars.Grammar
	execution  Execution
	parameters Parameters
}

func createElement(
	grammar grammars.Grammar,
	execution Execution,
	parameters Parameters,
) Element {
	out := element{
		grammar:    grammar,
		execution:  execution,
		parameters: parameters,
	}

	return &out
}

// Grammar returns the grammar
func (obj *element) Grammar() grammars.Grammar {
	return obj.grammar
}

// Execution returns the execution
func (obj *element) Execution() Execution {
	return obj.execution
}

// Parameters returns the parameters
func (obj *element) Parameters() Parameters {
	return obj.parameters
}
