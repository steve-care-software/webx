package layers

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/references"
)

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(instructions instructions.Instructions, output outputs.Output) Layer {
	ins, err := NewBuilder().Create().WithInstructions(instructions).WithOutput(output).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithReferencesForTests creates a new layer with references for tests
func NewLayerWithReferencesForTests(instructions instructions.Instructions, output outputs.Output, references references.References) Layer {
	ins, err := NewBuilder().Create().WithInstructions(instructions).WithOutput(output).WithReferences(references).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithInputForTests creates a new layer with input for tests
func NewLayerWithInputForTests(instructions instructions.Instructions, output outputs.Output, input string) Layer {
	ins, err := NewBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLayerWithReferencesAndInputForTests creates a new layer with references and input for tests
func NewLayerWithReferencesAndInputForTests(instructions instructions.Instructions, output outputs.Output, input string, references references.References) Layer {
	ins, err := NewBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).WithReferences(references).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
