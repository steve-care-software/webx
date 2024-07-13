package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/references"
)

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(instructions instructions.Instructions, output outputs.Output, input string) Layer {
	ins, err := NewBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithReferencesForTests creates a new layer with references for tests
func NewLayerWithReferencesForTests(instructions instructions.Instructions, output outputs.Output, input string, references references.References) Layer {
	ins, err := NewBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).WithReferences(references).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
