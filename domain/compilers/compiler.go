package compilers

type compiler struct {
	elements []Element
}

func createCompiler(
	elements []Element,
) Compiler {
	out := compiler{
		elements: elements,
	}

	return &out
}

// Elements returns the elements
func (obj *compiler) Elements() []Element {
	return obj.elements
}
