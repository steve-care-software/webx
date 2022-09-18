package grammars

type suite struct {
	valid   []byte
	invalid []byte
}

func createSuiteWithValid(
	valid []byte,
) Suite {
	return createSuiteInternally(valid, nil)
}

func createSuiteWithInvalid(
	invalid []byte,
) Suite {
	return createSuiteInternally(nil, invalid)
}

func createSuiteInternally(
	valid []byte,
	invalid []byte,
) Suite {
	out := suite{
		valid:   valid,
		invalid: invalid,
	}

	return &out
}

// IsValid returns true if valid, false otherwise
func (obj *suite) IsValid() bool {
	return obj.valid != nil
}

// Valid returns the valid bytes, if any
func (obj *suite) Valid() []byte {
	return obj.valid
}

// IsInvalid returns true if invalid, false otherwise
func (obj *suite) IsInvalid() bool {
	return obj.invalid != nil
}

// Invalid returns the invalid bytes, if any
func (obj *suite) Invalid() []byte {
	return obj.invalid
}
