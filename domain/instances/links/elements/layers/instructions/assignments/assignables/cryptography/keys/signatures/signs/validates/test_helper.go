package validates

// NewValidateForTests creates a new validate for tests
func NewValidateForTests(signature string, message string, publicKey string) Validate {
	ins, err := NewBuilder().Create().
		WithSignature(signature).
		WithMessage(message).
		WithPublicKey(publicKey).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
