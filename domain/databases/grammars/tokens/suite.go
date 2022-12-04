package tokens

type suite struct {
	isValid bool
	content []byte
}

func createSuite(
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

// Content returns the content
func (obj *suite) Content() []byte {
	return obj.content
}
