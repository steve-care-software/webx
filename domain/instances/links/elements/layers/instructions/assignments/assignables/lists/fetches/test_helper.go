package fetches

// NewFetchForTests creates a new fetch for tests
func NewFetchForTests(list string, index string) Fetch {
	ins, err := NewBuilder().Create().WithList(list).WithIndex(index).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
