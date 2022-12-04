package tokens

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
