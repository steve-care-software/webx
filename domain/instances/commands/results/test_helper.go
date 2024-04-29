package results

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs/kinds"
)

// NewResultWithInterruptionForTests creates a new result with interruption for tests
func NewResultWithInterruptionForTests(interruption interruptions.Interruption) Result {
	ins, err := NewBuilder().Create().WithInterruption(interruption).Now()
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
func NewSuccessForTests(output Output, kind kinds.Kind) Success {
	ins, err := NewSuccessBuilder().Create().WithOutput(output).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputForTests creates a new output for tests
func NewOutputForTests(input []byte) Output {
	ins, err := NewOutputBuilder().Create().WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputWithExecuteForTests creates a new output with execute for tests
func NewOutputWithExecuteForTests(input []byte, execute []byte) Output {
	ins, err := NewOutputBuilder().Create().WithInput(input).WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
