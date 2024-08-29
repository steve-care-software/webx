package instructions

type parameter struct {
	name  string
	value Value
}

func createParameter(
	name string,
	value Value,
) Parameter {
	out := parameter{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}

// Value returns the value
func (obj *parameter) Value() Value {
	return obj.value
}
