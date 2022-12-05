package programs

type program struct {
	instructions []Instruction
	outputs      [][]byte
}

func createProgram(
	instructions []Instruction,
) Program {
	return createProgramInternally(instructions, nil)
}

func createProgramWithOutputs(
	instructions []Instruction,
	outputs [][]byte,
) Program {
	return createProgramInternally(instructions, outputs)
}

func createProgramInternally(
	instructions []Instruction,
	outputs [][]byte,
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
func (obj *program) Outputs() [][]byte {
	return obj.outputs
}
