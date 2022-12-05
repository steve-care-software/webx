package compilers

type compiler struct {
	outputs  []string
	elements Elements
}

func createCompiler(
	elements Elements,
) Compiler {
	return createCompilerInternally(elements, nil)
}

func createCompilerWithOutputs(
	elements Elements,
	outputs []string,
) Compiler {
	return createCompilerInternally(elements, outputs)
}

func createCompilerInternally(
	elements Elements,
	outputs []string,
) Compiler {
	out := compiler{
		elements: elements,
		outputs:  outputs,
	}

	return &out
}

// Elements returns the elements
func (obj *compiler) Elements() Elements {
	return obj.elements
}

// HasOutputs returns true if there is outputs, false otherwise
func (obj *compiler) HasOutputs() bool {
	return obj.outputs != nil
}

// Outputs returns the outputs, if any
func (obj *compiler) Outputs() []string {
	return obj.outputs
}
