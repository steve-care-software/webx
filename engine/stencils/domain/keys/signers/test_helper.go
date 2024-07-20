package signers

// NewRingForTests creates a new ring for tests
func NewRingForTests(amount uint) []PublicKey {
	output := []PublicKey{}
	factory := NewFactory()
	casted := int(amount)
	for i := 0; i < casted; i++ {
		output = append(output, factory.Create().PublicKey())
	}

	return output
}
