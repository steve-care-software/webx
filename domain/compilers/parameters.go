package compilers

type parameters struct {
	list []Parameter
}

func createParameters(
	list []Parameter,
) Parameters {
	out := parameters{
		list: list,
	}

	return &out
}

// List returns the parameters
func (obj *parameters) List() []Parameter {
	return obj.list
}
