package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/outputs"
)

// NewLayersForTests creates a new layers for tests
func NewLayersForTests(list []Layer) Layers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(instructions instructions.Instructions, output outputs.Output, input string) Layer {
	ins, err := NewLayerBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
