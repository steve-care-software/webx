package executes

import "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"

// NewExecuteForTests creates a new execute for tests
func NewExecuteForTests(context string, input inputs.Input, ret string) Execute {
	ins, err := NewBuilder().Create().WithContext(context).WithInput(input).WithReturn(ret).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecuteWithLayerForTests creates a new execute with layer for tests
func NewExecuteWithLayerForTests(context string, input inputs.Input, ret string, layer inputs.Input) Execute {
	ins, err := NewBuilder().Create().WithContext(context).WithInput(input).WithReturn(ret).WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
