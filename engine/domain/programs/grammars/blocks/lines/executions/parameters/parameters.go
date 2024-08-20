package parameters

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

// List returns the list of parameter
func (obj *parameters) List() []Parameter {
	return obj.list
}
