package signs

// NewSignForTests creates a new sign for tests
func NewSignForTests(message string, account string) Sign {
	ins, err := NewBuilder().Create().
		WithMessage(message).
		WithAccount(account).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
