package begins

// NewBeginForTests creates a new begin for tests
func NewBeginForTests(path string, context string) Begin {
	ins, err := NewBuilder().Create().WithContext(context).WithPath(path).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
