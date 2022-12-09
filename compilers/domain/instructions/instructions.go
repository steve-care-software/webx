package instructions

type instructions struct {
	list      []Instruction
	remaining []byte
}

func createInstructions(
	list []Instruction,
) Instructions {
	return createInstructionsInternally(list, nil)
}

func createInstructionsWithRemaining(
	list []Instruction,
	remaining []byte,
) Instructions {
	return createInstructionsInternally(list, remaining)
}

func createInstructionsInternally(
	list []Instruction,
	remaining []byte,
) Instructions {
	out := instructions{
		list:      list,
		remaining: remaining,
	}

	return &out
}

// List returns the instructions
func (obj *instructions) List() []Instruction {
	return obj.list
}

// HasRemaining returns true if there is remaining data, false otherwise
func (obj *instructions) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining data, if any
func (obj *instructions) Remaining() []byte {
	return obj.remaining
}
