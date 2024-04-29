package deletes

// NewDeleteForTests creates a new delete for tests
func NewDeleteForTests(context string, path string, identifier string) Delete {
	ins, err := NewBuilder().Create().
		WithContext(context).
		WithPath(path).
		WithIdentifier(identifier).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
