package retrieves

// NewRetrieveForTests creates a new retrieve for tests
func NewRetrieveForTests(context string, index string) Retrieve {
	ins, err := NewBuilder().Create().WithContext(context).WithIndex(index).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRetrieveWithLengthForTests creates a new retrieve with length for tests
func NewRetrieveWithLengthForTests(context string, index string, length string) Retrieve {
	ins, err := NewBuilder().Create().WithContext(context).WithIndex(index).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
