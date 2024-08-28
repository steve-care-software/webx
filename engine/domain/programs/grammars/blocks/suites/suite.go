package suites

type suite struct {
	name   string
	value  []byte
	isFail bool
}

func createSuite(
	name string,
	value []byte,
	isFail bool,
) Suite {
	out := suite{
		name:   name,
		value:  value,
		isFail: isFail,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Value returns the value
func (obj *suite) Value() []byte {
	return obj.value
}

// IsFail returns true if expected to fail, false otherwise
func (obj *suite) IsFail() bool {
	return obj.isFail
}
