package inserts

// NewInsertForTests creates a new insert for tests
func NewInsertForTests(list string, element string) Insert {
	ins, err := NewBuilder().Create().WithList(list).WithElement(element).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
