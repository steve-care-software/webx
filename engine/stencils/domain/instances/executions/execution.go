package executions

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results"
	source_layers "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers"
)

type execution struct {
	hash   hash.Hash
	input  []byte
	source source_layers.Layer
	result results.Result
}

func createExecution(
	hash hash.Hash,
	input []byte,
	source source_layers.Layer,
	result results.Result,
) Execution {
	out := execution{
		hash:   hash,
		input:  input,
		source: source,
		result: result,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *execution) Input() []byte {
	return obj.input
}

// Source returns the source
func (obj *execution) Source() source_layers.Layer {
	return obj.source
}

// Result returns the result
func (obj *execution) Result() results.Result {
	return obj.result
}
