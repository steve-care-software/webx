package compilers

type compiler struct {
	outputs    []uint
	executions Executions
}

func createCompiler(
	executions Executions,
) Compiler {
	return createCompilerInternally(executions, nil)
}

func createCompilerWithOutputs(
	executions Executions,
	outputs []uint,
) Compiler {
	return createCompilerInternally(executions, outputs)
}

func createCompilerInternally(
	executions Executions,
	outputs []uint,
) Compiler {
	out := compiler{
		executions: executions,
		outputs:    outputs,
	}

	return &out
}

// Executions returns the executions
func (obj *compiler) Executions() Executions {
	return obj.executions
}

// HasOutputs returns true if there is outputs, false otherwise
func (obj *compiler) HasOutputs() bool {
	return obj.outputs != nil
}

// Outputs returns the outputs, if any
func (obj *compiler) Outputs() []uint {
	return obj.outputs
}
