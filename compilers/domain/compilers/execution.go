package compilers

import "github.com/steve-care-software/webx/programs/domain/programs"

type execution struct {
	parameters           Parameters
	program              programs.Program
	executeProgramModule uint
}

func createExecution(
	parameters Parameters,
	program programs.Program,
	executeProgramModule uint,
) Execution {
	out := execution{
		parameters:           parameters,
		program:              program,
		executeProgramModule: executeProgramModule,
	}

	return &out
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
