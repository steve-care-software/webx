package values

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
)

type value struct {
	parameter parameters.Parameter
	token     tokens.Token
}

func createValueWithParameter(
	parameter parameters.Parameter,
) Value {
	return createValueInternally(parameter, nil)
}

func createValueWithToken(
	token tokens.Token,
) Value {
	return createValueInternally(nil, token)
}

func createValueInternally(
	parameter parameters.Parameter,
	token tokens.Token,
) Value {
	out := value{
		parameter: parameter,
		token:     token,
	}

	return &out
}

// IsParameter returns true if there is a parameter, false otherwise
func (obj *value) IsParameter() bool {
	return obj.parameter != nil
}

// Parameter returns the parameter, if any
func (obj *value) Parameter() parameters.Parameter {
	return obj.parameter
}

// IsToken returns true if there is an token, false otherwise
func (obj *value) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *value) Token() tokens.Token {
	return obj.token
}
