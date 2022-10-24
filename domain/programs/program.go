package programs

type program struct {
	instructions []Instruction
	outputs      []string
}

func createProgram(
	instructions []Instruction,
) Program {
	return createProgramInternally(instructions, nil)
}

func createProgramWithOutputs(
	instructions []Instruction,
	outputs []string,
) Program {
	return createProgramInternally(instructions, outputs)
}

func createProgramInternally(
	instructions []Instruction,
	outputs []string,
) Program {
	out := program{
		instructions: instructions,
		outputs:      outputs,
	}

	return &out
}

// Instructions returns the instructions
func (obj *program) Instructions() []Instruction {
	return obj.instructions
}

// HasOutputs returns true if there is outputs, false otherwise
func (obj *program) HasOutputs() bool {
	return obj.outputs != nil
}

// Outputs returns the outputs, if any
func (obj *program) Outputs() []string {
	return obj.outputs
}
