package executions

import "github.com/steve-care-software/webx/engine/domain/grammars/tokens"

type execution struct {
	tokens tokens.Tokens
	fnName string
}

func createExecution(
	fnName string,
) Execution {
	return createExecutionInternally(fnName, nil)
}

func createExecutionWithTokens(
	fnName string,
	tokens tokens.Tokens,
) Execution {
	return createExecutionInternally(fnName, tokens)
}

func createExecutionInternally(
	fnName string,
	tokens tokens.Tokens,
) Execution {
	out := execution{
		fnName: fnName,
		tokens: tokens,
	}

	return &out
}

// FuncName returns the func name
func (obj *execution) FuncName() string {
	return obj.fnName
}

// HasTokens returns true if there is tokens, false otherwise
func (obj *execution) HasTokens() bool {
	return obj.tokens != nil
}

// Tokens returns the tokens
func (obj *execution) Tokens() tokens.Tokens {
	return obj.tokens
}
