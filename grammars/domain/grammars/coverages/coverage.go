package coverages

import "github.com/steve-care-software/webx/grammars/domain/grammars"

type coverage struct {
	token      grammars.Token
	executions Executions
}

func createCoverage(
	token grammars.Token,
	executions Executions,
) Coverage {
	out := coverage{
		token:      token,
		executions: executions,
	}

	return &out
}

// Token returns the token
func (obj *coverage) Token() grammars.Token {
	return obj.token
}

// Executions returns the executions
func (obj *coverage) Executions() Executions {
	return obj.executions
}
