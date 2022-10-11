package grammars

type suite struct {
	isValid bool
	content []byte
}

func createSuiteWithValid(
	valid []byte,
) Suite {
	return createSuiteInternally(true, valid)
}

func createSuiteWithInvalid(
	invalid []byte,
) Suite {
	return createSuiteInternally(false, invalid)
}

func createSuiteInternally(
	isValid bool,
	content []byte,
) Suite {
	out := suite{
		isValid: isValid,
		content: content,
	}

	return &out
}

// IsValid returns true if valid, false otherwise
func (obj *suite) IsValid() bool {
	return obj.isValid
}

// Content returns the the content
func (obj *suite) Content() []byte {
	return obj.content
}
