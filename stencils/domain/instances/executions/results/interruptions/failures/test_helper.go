package failures

// NewFailureForTests creates a new failure for tests
func NewFailureForTests(index uint, code uint, isRaisedInLayer bool) Failure {
	builder := NewBuilder().Create().WithIndex(index).WithCode(code)
	if isRaisedInLayer {
		builder.IsRaisedInLayer()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFailureWithMessageForTests creates a new failure with message for tests
func NewFailureWithMessageForTests(index uint, code uint, isRaisedInLayer bool, message string) Failure {
	builder := NewBuilder().Create().WithIndex(index).WithCode(code).WithMessage(message)
	if isRaisedInLayer {
		builder.IsRaisedInLayer()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
