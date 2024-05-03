package validates

// NewValidateForTests creates a new validate for tests
func NewValidateForTests(vote string, message string, hashedRing string) Validate {
	ins, err := NewBuilder().Create().WithVote(vote).WithMessage(message).WithHashedRing(hashedRing).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
