package executables

// NewExecutableWithLocalForTests creates a new Executable with local for tests
func NewExecutableWithLocalForTests(local string) Executable {
	ins, err := NewBuilder().Create().WithLocal(local).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutableWithRemoteForTests creates a new Executable with remote for tests
func NewExecutableWithRemoteForTests(remote string) Executable {
	ins, err := NewBuilder().Create().WithRemote(remote).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
