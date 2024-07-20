package results

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/success"
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
