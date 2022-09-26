package programs

import "github.com/steve-care-software/syntax/domain/syntax/programs/instructions"

type program struct {
	instructions instructions.Instructions
	inputs       []string
	outputs      []string
}

func createProgram(
	instructions instructions.Instructions,
) Program {
	return createProgramInternally(instructions, nil, nil)
}

func createProgramWithInputs(
	instructions instructions.Instructions,
	inputs []string,
) Program {
	return createProgramInternally(instructions, inputs, nil)
}

func createProgramWithOutputs(
	instructions instructions.Instructions,
	outputs []string,
) Program {
	return createProgramInternally(instructions, nil, outputs)
}

func createProgramWithInputsAndOutputs(
	instructions instructions.Instructions,
	inputs []string,
	outputs []string,
) Program {
	return createProgramInternally(instructions, inputs, outputs)
}

func createProgramInternally(
	instructions instructions.Instructions,
	inputs []string,
	outputs []string,
) Program {
	out := program{
		instructions: instructions,
		inputs:       inputs,
		outputs:      outputs,
	}

	return &out
}

// Instructions returns the instructions
func (obj *program) Instructions() instructions.Instructions {
	return obj.instructions
}

// HasInputs returns true if there is inputs, false otherwise
func (obj *program) HasInputs() bool {
	return obj.inputs != nil
}

// Inputs returns the inputs, if any
func (obj *program) Inputs() []string {
	return obj.inputs
}

// HasOutputs returns true if there is outputs, false otherwise
func (obj *program) HasOutputs() bool {
	return obj.outputs != nil
}

// Outputs returns the outputs, if any
func (obj *program) Outputs() []string {
	return obj.outputs
}
