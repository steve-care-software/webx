package votes

// NewVoteForTests creates a new vote for tests
func NewVoteForTests(message string, ring string, account string) Vote {
	ins, err := NewBuilder().Create().
		WithMessage(message).
		WithRing(ring).
		WithAccount(account).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
