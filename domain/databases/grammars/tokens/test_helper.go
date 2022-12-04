package tokens

// NewSuitesForTests creates a new suites for tests
func NewSuitesForTests(amount int) Suites {
	list := []Suite{}
	for i := 0; i < amount; i++ {
		isValid := i%2 == 0
		list = append(list, NewSuiteForTests(isValid))
	}

	ins, err := NewSuitesBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSuiteForTests creates a new suite for tests
func NewSuiteForTests(isValid bool) Suite {
	content := []byte("this is some content")
	builder := NewSuiteBuilder().Create().WithContent(content)
	if isValid {
		builder.IsValid()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
