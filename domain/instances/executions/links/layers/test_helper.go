package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/layers"
)

// NewLayersForTests creates new layers for tests
func NewLayersForTests(list []Layer) Layers {
	ins, err := NewBuilder().Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(input []byte, source source_layers.Layer, result results.Result) Layer {
	ins, err := NewLayerBuilder().Create().
		WithInput(input).
		WithSource(source).
		WithResult(result).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
