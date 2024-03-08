package layers

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/outputs/kinds"
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
func NewLayerForTests(instructions instructions.Instructions, output Output, input string) Layer {
	ins, err := NewLayerBuilder().Create().WithInstructions(instructions).WithOutput(output).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputWithExecuteForTests creates a new output with execute for tests
func NewOutputWithExecuteForTests(variable string, kind kinds.Kind, execute string) Output {
	ins, err := NewOutputBuilder().Create().WithVariable(variable).WithKind(kind).WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputForTests creates a new output for tests
func NewOutputForTests(variable string, kind kinds.Kind) Output {
	ins, err := NewOutputBuilder().Create().WithVariable(variable).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
