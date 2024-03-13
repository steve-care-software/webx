package inserts

// NewInsertForTests creates a new insert for tests
func NewInsertForTests(context string, instance string, path string) Insert {
	ins, err := NewBuilder().Create().
		WithContext(context).
		WithInstance(instance).
		WithPath(path).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
