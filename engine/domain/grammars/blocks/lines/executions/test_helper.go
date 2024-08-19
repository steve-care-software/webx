package executions

import "github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(fnName string) Execution {
	ins, err := NewBuilder().Create().WithFuncName(fnName).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithElementsForTests creates a new execution with elements for tests
func NewExecutionWithElementsForTests(elements elements.Elements, fnName string) Execution {
	ins, err := NewBuilder().Create().WithElements(elements).WithFuncName(fnName).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
