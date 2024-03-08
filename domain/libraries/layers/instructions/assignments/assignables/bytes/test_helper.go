package bytes

// NewBytesWithHashBytesForTests creates a new bytes with hashBytes for tests
func NewBytesWithHashBytesForTests(input string) Bytes {
	ins, err := NewBuilder().Create().WithHashBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesWithCompareForTests creates a new bytes with compare for tests
func NewBytesWithCompareForTests(input []string) Bytes {
	ins, err := NewBuilder().Create().WithCompare(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewBytesWithJoinForTests creates a new bytes with join for tests
func NewBytesWithJoinForTests(join []string) Bytes {
	ins, err := NewBuilder().Create().WithJoin(join).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
