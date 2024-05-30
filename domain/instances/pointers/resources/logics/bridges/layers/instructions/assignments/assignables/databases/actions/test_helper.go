package actions

// NewActionForTests creates a new action for tests
func NewActionForTests(path string, modifications string) Action {
	ins, err := NewBuilder().Create().WithPath(path).WithModifications(modifications).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
