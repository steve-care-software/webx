package outputs

import "github.com/steve-care-software/datastencil/states/domain/hash"

type output struct {
	hash    hash.Hash
	input   []byte
	execute []byte
}

func createOutput(
	hash hash.Hash,
	input []byte,
) Output {
	return createOutputInternally(hash, input, nil)
}

func createOutputWithExecute(
	hash hash.Hash,
	input []byte,
	execute []byte,
) Output {
	return createOutputInternally(hash, input, execute)
}

func createOutputInternally(
	hash hash.Hash,
	input []byte,
	execute []byte,
) Output {
	out := output{
		hash:    hash,
		input:   input,
		execute: execute,
	}

	return &out
}

// Hash retruns the hash
func (obj *output) Hash() hash.Hash {
	return obj.hash
}

// Value retruns the value
func (obj *output) Value() []byte {
	if obj.execute != nil {
		return obj.execute
	}

	return obj.input
}

// Input retruns the input
func (obj *output) Input() []byte {
	return obj.input
}

// HasExecute returns true if execute, false otherwise
func (obj *output) HasExecute() bool {
	return obj.execute != nil
}

// Execute returns the execute, if any
func (obj *output) Execute() []byte {
	return obj.execute
}
