package values

import (
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls/values/parameters"
)

type value struct {
	parameter parameters.Parameter
	token     string
}

func createValueWithParameter(
	parameter parameters.Parameter,
) Value {
	return createValueInternally(parameter, "")
}

func createValueWithToken(
	token string,
) Value {
	return createValueInternally(nil, token)
}

func createValueInternally(
	parameter parameters.Parameter,
	token string,
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
	return obj.token != ""
}

// Token returns the token, if any
func (obj *value) Token() string {
	return obj.token
}
