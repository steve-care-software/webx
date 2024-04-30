package integers

// NewIntegerWithIsEqualForTests creates a new integer operator with IsEqual for tests
func NewIntegerWithIsEqualForTests() Integer {
	ins, err := NewBuilder().Create().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerWithIsBiggerThanForTests creates a new integer operator with IsBiggerThan for tests
func NewIntegerWithIsBiggerThanForTests() Integer {
	ins, err := NewBuilder().Create().IsBiggerThan().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerWithIsBiggerThanAndIsEqualForTests creates a new integer operator with IsBiggerThan and isEqual for tests
func NewIntegerWithIsBiggerThanAndIsEqualForTests() Integer {
	ins, err := NewBuilder().Create().IsBiggerThan().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerWithIsSmallerThanForTests creates a new integer operator with IsSmallerThan for tests
func NewIntegerWithIsSmallerThanForTests() Integer {
	ins, err := NewBuilder().Create().IsSmallerThan().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerWithIsSmallerThanAndIsEqualForTests creates a new integer operator with IsSmallerThan and isEqual for tests
func NewIntegerWithIsSmallerThanAndIsEqualForTests() Integer {
	ins, err := NewBuilder().Create().IsSmallerThan().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
