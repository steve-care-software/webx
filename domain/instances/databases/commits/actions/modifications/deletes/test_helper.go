package deletes

// NewDeleteForTests creates a new delete for tests
func NewDeleteForTests(index uint, length uint) Delete {
	ins, err := NewBuilder().Create().WithIndex(index).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
