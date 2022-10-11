package grammars

type block struct {
	lines []Line
}

func createBlock(
	lines []Line,
) Block {
	out := block{
		lines: lines,
	}

	return &out
}

// Lines returns the lines
func (obj *block) Lines() []Line {
	return obj.lines
}
