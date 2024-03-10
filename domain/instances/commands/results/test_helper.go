package results

import "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs/kinds"

// NewResultWithFailureForTests creates a new result with failure for tests
func NewResultWithFailureForTests(failure Failure) Result {
	ins, err := NewBuilder().Create().WithFailure(failure).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResultWithSuccessForTests creates a new result with success for tests
func NewResultWithSuccessForTests(success Success) Result {
	ins, err := NewBuilder().Create().WithSuccess(success).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSuccessForTests creates a new success for tests
func NewSuccessForTests(bytes []byte, kind kinds.Kind) Success {
	ins, err := NewSuccessBuilder().Create().WithBytes(bytes).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFailureForTests creates a new failure for tests
func NewFailureForTests(index uint, code uint, isRaisedInLayer bool) Failure {
	builder := NewFailureBuilder().Create().WithIndex(index).WithCode(code)
	if isRaisedInLayer {
		builder.IsRaisedInLayer()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
