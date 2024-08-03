package executions

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results"
	source_layers "github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
)

type execution struct {
	hash   hash.Hash
	source source_layers.Layer
	result results.Result
	input  []byte
}

func createExecution(
	hash hash.Hash,
	source source_layers.Layer,
	result results.Result,
) Execution {
	return createExecutionInternally(hash, source, result, nil)
}

func createExecutionWithInput(
	hash hash.Hash,
	source source_layers.Layer,
	result results.Result,
	input []byte,
) Execution {
	return createExecutionInternally(hash, source, result, input)
}

func createExecutionInternally(
	hash hash.Hash,
	source source_layers.Layer,
	result results.Result,
	input []byte,
) Execution {
	out := execution{
		hash:   hash,
		source: source,
		result: result,
		input:  input,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Source returns the source
func (obj *execution) Source() source_layers.Layer {
	return obj.source
}

// Result returns the result
func (obj *execution) Result() results.Result {
	return obj.result
}

// HasInput returns true if there is an input, false otherwise
func (obj *execution) HasInput() bool {
	return obj.input != nil
}

// Input returns the input
func (obj *execution) Input() []byte {
	return obj.input
}
