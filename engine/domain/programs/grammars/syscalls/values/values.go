package values

type values struct {
	list []Value
}

func createValues(
	list []Value,
) Values {
	out := values{
		list: list,
	}

	return &out
}

// List returns the list of value
func (obj *values) List() []Value {
	return obj.list
}
