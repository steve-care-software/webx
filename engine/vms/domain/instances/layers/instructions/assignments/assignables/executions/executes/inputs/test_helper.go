package inputs

// NewInputWithValueForTests creates a new input with value for tests
func NewInputWithValueForTests(value string) Input {
	ins, err := NewBuilder().Create().WithValue(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInputWithPathForTests creates a new input with path for tests
func NewInputWithPathForTests(path string) Input {
	ins, err := NewBuilder().Create().WithPath(path).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
