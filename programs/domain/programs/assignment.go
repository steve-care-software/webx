package programs

type assignment struct {
	name  []byte
	value Value
}

func createAssignment(
	name []byte,
	value Value,
) Assignment {
	out := assignment{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *assignment) Name() []byte {
	return obj.name
}

// Value returns the value
func (obj *assignment) Value() Value {
	return obj.value
}
