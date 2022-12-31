package compilers

import (
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/programs/domain/programs"
)

type execution struct {
	grammar              grammars.Grammar
	parameters           Parameters
	program              programs.Program
	executeProgramModule uint
}

func createExecution(
	grammar grammars.Grammar,
	parameters Parameters,
	program programs.Program,
	executeProgramModule uint,
) Execution {
	out := execution{
		grammar:              grammar,
		parameters:           parameters,
		program:              program,
		executeProgramModule: executeProgramModule,
	}

	return &out
}

// Grammar returns the grammar
func (obj *execution) Grammar() grammars.Grammar {
	return obj.grammar
}

// Parameters returns the parameters
func (obj *execution) Parameters() Parameters {
	return obj.parameters
}

// Program returns the program
func (obj *execution) Program() programs.Program {
	return obj.program
}

// ExecuteProgramModule returns the execute program module
func (obj *execution) ExecuteProgramModule() uint {
	return obj.executeProgramModule
}
