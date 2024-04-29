package decrypts

// NewDecryptForTests creates a new decrypt for tests
func NewDecryptForTests(cipher string, account string) Decrypt {
	ins, err := NewBuilder().Create().WithCipher(cipher).WithAccount(account).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
