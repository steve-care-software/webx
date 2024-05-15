package deletes

// NewDeleteForTests creates a new delete for tests
func NewDeleteForTests(list string, index string) Delete {
	ins, err := NewBuilder().Create().WithList(list).WithIndex(index).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
