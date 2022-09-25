package outputs

type variable struct {
	name  string
	value interface{}
}

func createVariable(
	name string,
) Variable {
	return createVariableInternally(name, nil)
}

func createVariableWithValue(
	name string,
	value interface{},
) Variable {
	return createVariableInternally(name, value)
}

func createVariableInternally(
	name string,
	value interface{},
) Variable {
	out := variable{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *variable) Name() string {
	return obj.name
}

// HasValue returns true if there is value, false otherwise
func (obj *variable) HasValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *variable) Value() interface{} {
	return obj.value
}
