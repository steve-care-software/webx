package relationals

// NewRelationalWithOrForTests creates a relational operator with or for tests
func NewRelationalWithOrForTests() Relational {
	ins, err := NewBuilder().Create().IsOr().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRelationalWithAndForTests creates a relational operator with and for tests
func NewRelationalWithAndForTests() Relational {
	ins, err := NewBuilder().Create().IsAnd().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
