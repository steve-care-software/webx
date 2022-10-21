package instructions

type instructions struct {
	list []Instruction
}

func createInstructions(
	list []Instruction,
) Instructions {
	out := instructions{
		list: list,
	}

	return &out
}

// List returns the instructions
func (obj *instructions) List() []Instruction {
	return obj.list
}
