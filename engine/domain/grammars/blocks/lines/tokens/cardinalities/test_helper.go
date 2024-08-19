package cardinalities

// NewCardinalityForTests creates a new cardinality for tests
func NewCardinalityForTests(min uint) Cardinality {
	ins, err := NewBuilder().Create().WithMin(min).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCardinalityWithMaxForTests creates a new cardinality with max for tests
func NewCardinalityWithMaxForTests(min uint, max uint) Cardinality {
	ins, err := NewBuilder().Create().WithMin(min).WithMax(max).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
