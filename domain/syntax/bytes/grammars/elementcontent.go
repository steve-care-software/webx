package grammars

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/values"
)

type elementContent struct {
	value    values.Value
	external External
	instance Instance
}

func createElementContentWithValue(
	value values.Value,
) ElementContent {
	return createElementContentInternally(value, nil, nil)
}

func createElementContentWithExternalToken(
	external External,
) ElementContent {
	return createElementContentInternally(nil, external, nil)
}

func createElementContentWithInstance(
	instance Instance,
) ElementContent {
	return createElementContentInternally(nil, nil, instance)
}

func createElementContentInternally(
	value values.Value,
	external External,
	instance Instance,
) ElementContent {
	out := elementContent{
		value:    value,
		external: external,
		instance: instance,
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

// IsExternal returns true if there is an external grammar, false otherwise
func (obj *elementContent) IsExternal() bool {
	return obj.external != nil
}

// External returns the external grammar, if any
func (obj *elementContent) External() External {
	return obj.external
}

// IsInstance returns true if there is an instance, false otherwise
func (obj *elementContent) IsInstance() bool {
	return obj.instance != nil
}

// Instance returns the instance, if any
func (obj *elementContent) Instance() Instance {
	return obj.instance
}
