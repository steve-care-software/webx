package decrypts

// NewDecryptForTests creates a new decrypt for tests
func NewDecryptForTests(cipher string, password string) Decrypt {
	ins, err := NewBuilder().Create().WithCipher(cipher).WithPassword(password).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
