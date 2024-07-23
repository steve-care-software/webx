package interruptions

import "github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/interruptions/failures"

// NewInterruptionWithStopForTests creates a new interruption with stop for tests
func NewInterruptionWithStopForTests(stop uint) Interruption {
	ins, err := NewBuilder().Create().WithStop(stop).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewInterruptionWithFailureForTests creates a new failure  with stop for tests
func NewInterruptionWithFailureForTests(failure failures.Failure) Interruption {
	ins, err := NewBuilder().Create().WithFailure(failure).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
