package profiles

// NewProfileForTests creates a new profile for tests
func NewProfileForTests(name string, description string) Profile {
	ins, err := NewBuilder().Create().WithName(name).WithDescription(description).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewProfileWithPackagesForTests creates a new profile with packages for tests
func NewProfileWithPackagesForTests(name string, description string) Profile {
	ins, err := NewBuilder().Create().WithName(name).WithDescription(description).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
