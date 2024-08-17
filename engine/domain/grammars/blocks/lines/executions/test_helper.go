package executions

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(flag uint16) Execution {
	ins, err := NewBuilder().Create().WithFuncFlag(flag).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithTokensForTests creates a new execution with tokens for tests
func NewExecutionWithTokensForTests(tokens []string, flag uint16) Execution {
	ins, err := NewBuilder().Create().WithTokens(tokens).WithFuncFlag(flag).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
