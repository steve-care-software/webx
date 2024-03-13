package reverts

// NewRevertForTests creates a new revert for tests
func NewRevertForTests() Revert {
	ins, err := NewBuilder().Create().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRevertWithIndexForTests creates a new revert with index for tests
func NewRevertWithIndexForTests(index string) Revert {
	ins, err := NewBuilder().Create().WithIndex(index).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
