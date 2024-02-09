package layers

import "github.com/steve-care-software/identity/domain/hash"

type execution struct {
	hash  hash.Hash
	input string
	layer string
}

func createExecution(
	hash hash.Hash,
	input string,
) Execution {
	return createExecutionInternally(hash, input, "")
}

func createExecutionWithLayer(
	hash hash.Hash,
	input string,
	layer string,
) Execution {
	return createExecutionInternally(hash, input, layer)
}

func createExecutionInternally(
	hash hash.Hash,
	input string,
	layer string,
) Execution {
	out := execution{
		hash:  hash,
		input: input,
		layer: layer,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *execution) Input() string {
	return obj.input
}

// HasLayer retruns true if there is layer, false otherwise
func (obj *execution) HasLayer() bool {
	return obj.layer != ""
}

// Layer returns the layer, if any
func (obj *execution) Layer() string {
	return obj.layer
}
