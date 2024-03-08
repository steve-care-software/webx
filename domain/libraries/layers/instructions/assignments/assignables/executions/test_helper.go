package executions

// NewExecutionWithLayerForTests creates a new execution with layer for tests
func NewExecutionWithLayerForTests(input string, layer string) Execution {
	ins, err := NewBuilder().Create().WithInput(input).WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(input string) Execution {
	ins, err := NewBuilder().Create().WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
