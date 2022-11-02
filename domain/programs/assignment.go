package programs

type assignment struct {
	index uint
	name  []byte
	value Value
}

func createAssignment(
	index uint,
	name []byte,
	value Value,
) Assignment {
	out := assignment{
		index: index,
		name:  name,
		value: value,
	}

	return &out
}

// Index returns the index
func (obj *assignment) Index() uint {
	return obj.index
}

// Name returns the name
func (obj *assignment) Name() []byte {
	return obj.name
}

// Value returns the value
func (obj *assignment) Value() Value {
	return obj.value
}
