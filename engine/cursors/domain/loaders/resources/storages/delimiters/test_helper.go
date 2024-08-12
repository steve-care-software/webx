package delimiters

// NewDelimitersForTests creates a new delimiters for tests
func NewDelimitersForTests(list []Delimiter) Delimiters {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDelimiterForTests creates a new delimiter for tests
func NewDelimiterForTests(index uint64, length uint64) Delimiter {
	ins, err := NewDelimiterBuilder().Create().WithIndex(index).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
