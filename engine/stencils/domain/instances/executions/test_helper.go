package executions

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results"
	source_layers "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers"
)

// NewExecutionsForTests creates a new executions for tests
func NewExecutionsForTests(list []Execution) Executions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(input []byte, source source_layers.Layer, result results.Result) Execution {
	ins, err := NewExecutionBuilder().Create().
		WithInput(input).
		WithSource(source).
		WithResult(result).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
