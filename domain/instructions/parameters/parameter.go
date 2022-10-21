package parameters

type parameter struct {
	name    string
	isInput bool
}

func createParameterWithInput(
	name string,
) Parameter {
	return createParameterInternally(name, true)
}

func createParameterWithOutput(
	name string,
) Parameter {
	return createParameterInternally(name, false)
}

func createParameterInternally(
	name string,
	isInput bool,
) Parameter {
	out := parameter{
		name:    name,
		isInput: isInput,
	}

	return &out
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}

// IsInput returns true if the parameter is an input, false otherwise
func (obj *parameter) IsInput() bool {
	return obj.isInput
}
