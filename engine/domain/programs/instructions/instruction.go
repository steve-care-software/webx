package instructions

import "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"

type instruction struct {
	block  string
	line   uint
	tokens tokens.Tokens
}

func createInstruction(
	block string,
	line uint,
	tokens tokens.Tokens,
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
func (obj *instruction) Tokens() tokens.Tokens {
	return obj.tokens
}
