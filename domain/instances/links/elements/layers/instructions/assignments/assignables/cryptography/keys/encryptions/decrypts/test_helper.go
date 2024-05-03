package decrypts

// NewDecryptForTests creates a new decrypt for tests
func NewDecryptForTests(cipher string, pk string) Decrypt {
	ins, err := NewBuilder().Create().
		WithCipher(cipher).
		WithPrivateKey(pk).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
