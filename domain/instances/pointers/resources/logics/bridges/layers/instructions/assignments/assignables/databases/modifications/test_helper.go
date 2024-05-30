package modifications

// NewModificationWithInsertForTests cretes a new modification with insert for tests
func NewModificationWithInsertForTests(insert string) Modification {
	ins, err := NewBuilder().Create().WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewModificationWithDeleteForTests cretes a new modification with insert for tests
func NewModificationWithDeleteForTests(delete string) Modification {
	ins, err := NewBuilder().Create().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
