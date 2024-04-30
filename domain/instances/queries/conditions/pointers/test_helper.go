package pointers

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(entity string, field string) Pointer {
	ins, err := NewBuilder().Create().WithEntity(entity).WithField(field).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
