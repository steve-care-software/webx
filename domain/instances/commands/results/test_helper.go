package results

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs/kinds"
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
func NewSuccessForTests(bytes []byte, kind kinds.Kind) Success {
	ins, err := NewSuccessBuilder().Create().WithBytes(bytes).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
