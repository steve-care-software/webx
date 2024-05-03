package encrypts

// NewEncryptForTests creates a new encrypt for tests
func NewEncryptForTests(message string, pubKey string) Encrypt {
	ins, err := NewBuilder().Create().WithMessage(message).WithPublicKey(pubKey).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
