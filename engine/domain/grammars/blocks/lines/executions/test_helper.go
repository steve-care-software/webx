package executions

import "github.com/steve-care-software/webx/engine/domain/grammars/tokens"

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(fnName string) Execution {
	ins, err := NewBuilder().Create().WithFuncName(fnName).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithTokensForTests creates a new execution with tokens for tests
func NewExecutionWithTokensForTests(tokens tokens.Tokens, fnName string) Execution {
	ins, err := NewBuilder().Create().WithTokens(tokens).WithFuncName(fnName).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
