package executes

import "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"

// NewExecuteForTests creates a new execute for tests
func NewExecuteForTests(context string, input inputs.Input) Execute {
	ins, err := NewBuilder().Create().WithContext(context).WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecuteWithLayerForTests creates a new execute with layer for tests
func NewExecuteWithLayerForTests(context string, input inputs.Input, layer string) Execute {
	ins, err := NewBuilder().Create().WithContext(context).WithInput(input).WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
