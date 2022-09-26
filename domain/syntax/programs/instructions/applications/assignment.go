package applications

type assignment struct {
	name  string
	value Value
}

func createAssignment(
	name string,
	value Value,
) Assignment {
	out := assignment{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *assignment) Name() string {
	return obj.name
}

// Value returns the value
func (obj *assignment) Value() Value {
	return obj.value
}
