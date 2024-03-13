package constants

// NewConstantWithBoolForTests creates constant with bool for tests
func NewConstantWithBoolForTests(value bool) Constant {
	ins, err := NewBuilder().Create().WithBool(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConstantWithBytesForTests creates constant with bytes for tests
func NewConstantWithBytesForTests(value []byte) Constant {
	ins, err := NewBuilder().Create().WithBytes(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
