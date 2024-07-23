package constants

// NewConstantsForTests creates a new constants for tests
func NewConstantsForTests(list []Constant) Constants {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConstantWithBytesForTests creates a new constant with bytes for tests
func NewConstantWithBytesForTests(value []byte) Constant {
	ins, err := NewConstantBuilder().Create().WithBytes(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConstantWithBoolForTests creates a new constant with bool for tests
func NewConstantWithBoolForTests(value bool) Constant {
	ins, err := NewConstantBuilder().Create().WithBool(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConstantWithStringForTests creates a new constant with string for tests
func NewConstantWithStringForTests(value string) Constant {
	ins, err := NewConstantBuilder().Create().WithString(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConstantWithIntForTests creates a new constant with int for tests
func NewConstantWithIntForTests(value int) Constant {
	ins, err := NewConstantBuilder().Create().WithInt(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConstantWithUintForTests creates a new constant with uint for tests
func NewConstantWithUintForTests(value uint) Constant {
	ins, err := NewConstantBuilder().Create().WithUint(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConstantWithFloatForTests creates a new constant with float for tests
func NewConstantWithFloatForTests(value float64) Constant {
	ins, err := NewConstantBuilder().Create().WithFloat(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConstantWithListForTests creates a new constant with list for tests
func NewConstantWithListForTests(value Constants) Constant {
	ins, err := NewConstantBuilder().Create().WithList(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
