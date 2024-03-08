package layers

import "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions"

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
func NewOutputWithExecuteForTests(variable string, kind Kind, execute string) Output {
	ins, err := NewOutputBuilder().Create().WithVariable(variable).WithKind(kind).WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputForTests creates a new output for tests
func NewOutputForTests(variable string, kind Kind) Output {
	ins, err := NewOutputBuilder().Create().WithVariable(variable).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewKindWithContinueForTests creates a new kind with continue for tests
func NewKindWithContinueForTests() Kind {
	ins, err := NewKindBuilder().Create().IsContinue().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewKindWithPromptForTests creates a new kind with prompt for tests
func NewKindWithPromptForTests() Kind {
	ins, err := NewKindBuilder().Create().IsPrompt().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
