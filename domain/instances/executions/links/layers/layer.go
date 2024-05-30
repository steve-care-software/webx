package layers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
)

type layer struct {
	hash   hash.Hash
	input  []byte
	source source_layers.Layer
	result results.Result
}

func createLayer(
	hash hash.Hash,
	input []byte,
	source source_layers.Layer,
	result results.Result,
) Layer {
	out := layer{
		hash:   hash,
		input:  input,
		source: source,
		result: result,
	}

	return &out
}

// Hash returns the hash
func (obj *layer) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *layer) Input() []byte {
	return obj.input
}

// Source returns the source
func (obj *layer) Source() source_layers.Layer {
	return obj.source
}

// Result returns the result
func (obj *layer) Result() results.Result {
	return obj.result
}
