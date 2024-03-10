package operators

// NewOperatorWithXOrForTests creates a new operator with xor for tests
func NewOperatorWithXOrForTests() Operator {
	ins, err := NewBuilder().Create().IsXor().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOperatorWithOrForTests creates a new operator with or for tests
func NewOperatorWithOrForTests() Operator {
	ins, err := NewBuilder().Create().IsOr().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOperatorWithAndForTests creates a new operator with and for tests
func NewOperatorWithAndForTests() Operator {
	ins, err := NewBuilder().Create().IsAnd().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
