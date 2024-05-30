package creates

// NewCreateForTests creates a new create for tests
func NewCreateForTests(message string, ring string, privateKey string) Create {
	ins, err := NewBuilder().Create().WithMessage(message).WithRing(ring).WithPrivateKey(privateKey).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
