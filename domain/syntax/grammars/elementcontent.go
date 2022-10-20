package grammars

import (
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/values"
)

type elementContent struct {
	hash      hash.Hash
	value     values.Value
	external  External
	instance  Instance
	recursive string
}

func createElementContentWithValue(
	hash hash.Hash,
	value values.Value,
) ElementContent {
	return createElementContentInternally(hash, value, nil, nil, "")
}

func createElementContentWithExternalToken(
	hash hash.Hash,
	external External,
) ElementContent {
	return createElementContentInternally(hash, nil, external, nil, "")
}

func createElementContentWithInstance(
	hash hash.Hash,
	instance Instance,
) ElementContent {
	return createElementContentInternally(hash, nil, nil, instance, "")
}

func createElementContentWithRecursive(
	hash hash.Hash,
	recursive string,
) ElementContent {
	return createElementContentInternally(hash, nil, nil, nil, recursive)
}

func createElementContentInternally(
	hash hash.Hash,
	value values.Value,
	external External,
	instance Instance,
	recursive string,
) ElementContent {
	out := elementContent{
		hash:      hash,
		value:     value,
		external:  external,
		instance:  instance,
		recursive: recursive,
	}

	return &out
}

// Hash returns the hash
func (obj *elementContent) Hash() hash.Hash {
	return obj.hash
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

// IsRecursive returns true if there is a recursive token, false otherwise
func (obj *elementContent) IsRecursive() bool {
	return obj.recursive != ""
}

// Recursive returns the recursive, if any
func (obj *elementContent) Recursive() string {
	return obj.recursive
}
