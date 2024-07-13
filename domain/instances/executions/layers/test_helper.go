package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/layers"
)

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(input []byte, source source_layers.Layer, result results.Result) Layer {
	ins, err := NewBuilder().Create().
		WithInput(input).
		WithSource(source).
		WithResult(result).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
