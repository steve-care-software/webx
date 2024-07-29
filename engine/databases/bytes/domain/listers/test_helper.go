package listers

// NewListerForTests creates a new lister for tests
func NewListerForTests(keyname string, index uint64, length uint64) Lister {
	ins, err := NewBuilder().Create().WithKeyname(keyname).WithIndex(index).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
