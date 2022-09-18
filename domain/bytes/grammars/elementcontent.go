package grammars

import (
	"github.com/steve-care-software/logics/domain/bytes/grammars/values"
)

type elementContent struct {
	value    values.Value
	token    Token
	external External
}

func createElementContentWithValue(
	value values.Value,
) ElementContent {
	return createElementContentInternally(value, nil, nil)
}

func createElementContentWithToken(
	token Token,
) ElementContent {
	return createElementContentInternally(nil, token, nil)
}

func createElementContentWithExternalToken(
	external External,
) ElementContent {
	return createElementContentInternally(nil, nil, external)
}

func createElementContentInternally(
	value values.Value,
	token Token,
	external External,
) ElementContent {
	out := elementContent{
		value:    value,
		token:    token,
		external: external,
	}

	return &out
}

// IsValue returns true if there is a value, false otherwise
func (obj *elementContent) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *elementContent) Value() values.Value {
	return obj.value
}

// IsToken returns true if there is a token, false otherwise
func (obj *elementContent) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *elementContent) Token() Token {
	return obj.token
}

// IsExternal returns true if there is an external token, false otherwise
func (obj *elementContent) IsExternal() bool {
	return obj.external != nil
}

// External returns the external token, if any
func (obj *elementContent) External() External {
	return obj.external
}
