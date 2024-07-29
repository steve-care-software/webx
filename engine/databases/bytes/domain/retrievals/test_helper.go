package retrievals

// NewRetrievalsForTests creates a new retrievals for tests
func NewRetrievalsForTests(list []Retrieval) Retrievals {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRetrievalForTests creates a new retrieval for tests
func NewRetrievalForTests(index uint64, length uint64) Retrieval {
	ins, err := NewRetrievalBuilder().Create().WithIndex(index).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
