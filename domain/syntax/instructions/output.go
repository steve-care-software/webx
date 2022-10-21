package instructions

type output struct {
	instructions Instructions
	remaining    []byte
}

func createOutput(
	instructions Instructions,
) Output {
	return createOutputInternally(instructions, nil)
}

func createOutputWithRemaining(
	instructions Instructions,
	remaining []byte,
) Output {
	return createOutputInternally(instructions, remaining)
}

func createOutputInternally(
	instructions Instructions,
	remaining []byte,
) Output {
	out := output{
		instructions: instructions,
		remaining:    remaining,
	}

	return &out
}

// Instructions returns the instructions
func (obj *output) Instructions() Instructions {
	return obj.instructions
}

// HasRemaining returns true if there is remaining, false otherwise
func (obj *output) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining, if any
func (obj *output) Remaining() []byte {
	return obj.remaining
}
