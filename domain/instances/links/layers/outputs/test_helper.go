package outputs

import "github.com/steve-care-software/datastencil/domain/instances/links/layers/outputs/kinds"

// NewOutputWithExecuteForTests creates a new output with execute for tests
func NewOutputWithExecuteForTests(variable string, kind kinds.Kind, execute []string) Output {
	ins, err := NewBuilder().Create().WithVariable(variable).WithKind(kind).WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputForTests creates a new output for tests
func NewOutputForTests(variable string, kind kinds.Kind) Output {
	ins, err := NewBuilder().Create().WithVariable(variable).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
