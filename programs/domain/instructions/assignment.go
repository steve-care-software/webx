package instructions

type assignment struct {
	variable []byte
	value    Value
}

func createAssignment(
	variable []byte,
	value Value,
) Assignment {
	out := assignment{
		variable: variable,
		value:    value,
	}

	return &out
}

// Variable returns the variable
func (obj *assignment) Variable() []byte {
	return obj.variable
}

// Value returns the value
func (obj *assignment) Value() Value {
	return obj.value
}
