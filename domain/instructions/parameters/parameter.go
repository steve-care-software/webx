package parameters

type parameter struct {
	name    []byte
	isInput bool
}

func createParameterWithInput(
	name []byte,
) Parameter {
	return createParameterInternally(name, true)
}

func createParameterWithOutput(
	name []byte,
) Parameter {
	return createParameterInternally(name, false)
}

func createParameterInternally(
	name []byte,
	isInput bool,
) Parameter {
	out := parameter{
		name:    name,
		isInput: isInput,
	}

	return &out
}

// Name returns the name
func (obj *parameter) Name() []byte {
	return obj.name
}

// IsInput returns true if the parameter is an input, false otherwise
func (obj *parameter) IsInput() bool {
	return obj.isInput
}
