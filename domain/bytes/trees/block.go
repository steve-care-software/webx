package trees

type block struct {
	lines      []Line
	successful Line
}

func createBlock(
	lines []Line,
) Block {
	return createBlockInternally(lines, nil)
}

func createBlockWithSuccessful(
	lines []Line,
	successful Line,
) Block {
	return createBlockInternally(lines, successful)
}

func createBlockInternally(
	lines []Line,
	successful Line,
) Block {
	out := block{
		lines:      lines,
		successful: successful,
	}

	return &out
}

// Lines returns the lines
func (obj *block) Lines() []Line {
	return obj.lines
}

// HasSuccessful returns true if there is a successful line, false otherwise
func (obj *block) HasSuccessful() bool {
	return obj.successful != nil
}

// Successful returns the successful line, if any
func (obj *block) Successful() Line {
	return obj.successful
}
