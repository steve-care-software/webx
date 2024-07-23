package opens

// NewOpenForTests creates a new open for tests
func NewOpenForTests(path string, permission string) Open {
	ins, err := NewBuilder().Create().WithPath(path).WithPermission(permission).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
