package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs"
)

// NewLayerForTests creates a new layer for tests
func NewLayerForTests(instructions instructions.Instructions, output outputs.Output, input string) Layer {
	ins, err := NewBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
