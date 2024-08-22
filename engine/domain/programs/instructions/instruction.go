package instructions

type instruction struct {
	block  string
	line   uint
	tokens Tokens
}

func createInstruction(
	block string,
	line uint,
	tokens Tokens,
) Instruction {
	out := instruction{
		block:  block,
		line:   line,
		tokens: tokens,
	}

	return &out
}

// Block returns the block
func (obj *instruction) Block() string {
	return obj.block
}

// Line returns the line
func (obj *instruction) Line() uint {
	return obj.line
}

// Line returns the line
func (obj *instruction) Tokens() Tokens {
	return obj.tokens
}
