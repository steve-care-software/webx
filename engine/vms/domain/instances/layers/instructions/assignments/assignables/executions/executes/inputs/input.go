package inputs

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type input struct {
	hash  hash.Hash
	value string
	path  string
}

func createInputWithValue(
	hash hash.Hash,
	value string,
) Input {
	return createInputInternally(hash, value, "")
}

func createInputWithPath(
	hash hash.Hash,
	path string,
) Input {
	return createInputInternally(hash, "", path)
}

func createInputInternally(
	hash hash.Hash,
	value string,
	path string,
) Input {
	out := input{
		hash:  hash,
		value: value,
		path:  path,
	}

	return &out
}

// Hash returns the hash
func (obj *input) Hash() hash.Hash {
	return obj.hash
}

// IsValue returns true if there is a value, false otherwise
func (obj *input) IsValue() bool {
	return obj.value != ""
}

// Value returns the value, if any
func (obj *input) Value() string {
	return obj.value
}

// IsPath returns true if there is a path, false otherwise
func (obj *input) IsPath() bool {
	return obj.path != ""
}

// Path returns the path, if any
func (obj *input) Path() string {
	return obj.path
}
