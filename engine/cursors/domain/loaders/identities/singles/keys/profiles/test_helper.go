package profiles

// NewProfileForTests creates a new profile for tests
func NewProfileForTests(name string, description string) Profile {
	ins, err := NewBuilder().Create().WithName(name).WithDescription(description).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewProfileWithNamespacesForTests creates a new profile with namespaces for tests
func NewProfileWithNamespacesForTests(name string, description string) Profile {
	ins, err := NewBuilder().Create().WithName(name).WithDescription(description).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
