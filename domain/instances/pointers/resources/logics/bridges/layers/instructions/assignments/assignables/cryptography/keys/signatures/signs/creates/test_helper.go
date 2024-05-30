package creates

// NewCreateForTests creates a new create for tests
func NewCreateForTests(message string, pk string) Create {
	ins, err := NewBuilder().Create().
		WithMessage(message).
		WithPrivateKey(pk).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
