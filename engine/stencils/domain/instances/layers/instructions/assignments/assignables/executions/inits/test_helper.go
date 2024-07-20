package inits

// NewInitForTests creates a new init for tests
func NewInitForTests(path string, name string, description string) Init {
	ins, err := NewBuilder().Create().WithPath(path).WithName(name).WithDescription(description).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
