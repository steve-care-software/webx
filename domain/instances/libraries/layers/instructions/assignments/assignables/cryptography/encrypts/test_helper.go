package encrypts

// NewEncryptForTests creates a new encrypt for tests
func NewEncryptForTests(message string, password string) Encrypt {
	ins, err := NewBuilder().Create().WithMessage(message).WithPassword(password).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
