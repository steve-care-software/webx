package retrieves

// NewRetrieveWithListForTests creates a new retrieve with list for tests
func NewRetrieveWithListForTests() Retrieve {
	ins, err := NewBuilder().Create().IsList().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRetrieveWithExistsForTests creates a new retrieve with exists for tests
func NewRetrieveWithExistsForTests(exists string) Retrieve {
	ins, err := NewBuilder().Create().WithExists(exists).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRetrieveWithRetrieveForTests creates a new retrieve with retrieve for tests
func NewRetrieveWithRetrieveForTests(retrieve string) Retrieve {
	ins, err := NewBuilder().Create().WithRetrieve(retrieve).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
