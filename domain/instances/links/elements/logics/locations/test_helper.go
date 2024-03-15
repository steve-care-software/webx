package locations

// NewLocationWithSingleForTests creates a new location with single for tests
func NewLocationWithSingleForTests(single []byte) Location {
	ins, err := NewBuilder().Create().WithSingle(single).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLocationWithListForTests creates a new location with list for tests
func NewLocationWithListForTests(list [][]byte) Location {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
