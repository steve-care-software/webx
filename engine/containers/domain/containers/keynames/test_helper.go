package keynames

// NewKeynameForTests creates a new keyname for tests
func NewKeynameForTests(name string) Keyname {
	ins, err := NewBuilder().Create().WithName(name).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
