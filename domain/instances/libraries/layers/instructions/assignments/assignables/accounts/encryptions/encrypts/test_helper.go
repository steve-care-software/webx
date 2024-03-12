package encrypts

// NewEncryptForTests creates a new encrypt for tests
func NewEncryptForTests(message string, account string) Encrypt {
	ins, err := NewBuilder().Create().WithMessage(message).WithAccount(account).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
