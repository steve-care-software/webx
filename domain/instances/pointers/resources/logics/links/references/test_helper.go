package references

// NewReferencesForTests creates references for tests
func NewReferencesForTests(list []Reference) References {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReferenceForTests creates a new reference for tests
func NewReferenceForTests(variable string, path []string) Reference {
	ins, err := NewReferenceBuilder().Create().
		WithVariable(variable).
		WithPath(path).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
