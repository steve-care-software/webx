package executions

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
)

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(fnName string) Execution {
	ins, err := NewBuilder().Create().WithFuncName(fnName).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithParameterssForTests creates a new execution with parameters for tests
func NewExecutionWithParameterssForTests(parameters parameters.Parameters, fnName string) Execution {
	ins, err := NewBuilder().Create().WithParameters(parameters).WithFuncName(fnName).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
