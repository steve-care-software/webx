package kinds

// NewKindWithContinueForTests creates a new kind with continue for tests
func NewKindWithContinueForTests() Kind {
	ins, err := NewBuilder().Create().IsContinue().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewKindWithPromptForTests creates a new kind with prompt for tests
func NewKindWithPromptForTests() Kind {
	ins, err := NewBuilder().Create().IsPrompt().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
