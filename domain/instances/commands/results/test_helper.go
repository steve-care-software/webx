package results

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/success"
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
func NewResultWithSuccessForTests(success success.Success) Result {
	ins, err := NewBuilder().Create().WithSuccess(success).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
