package trees

import "github.com/steve-care-software/syntax/domain/bytes/grammars"

type tree struct {
	grammar grammars.Token
	block   Block
}

func createTree(
	grammar grammars.Token,
	block Block,
) Tree {
	out := tree{
		grammar: grammar,
		block:   block,
	}

	return &out
}

// Grammar returns the grammar
func (obj *tree) Grammar() grammars.Token {
	return obj.grammar
}

// Block returns the block
func (obj *tree) Block() Block {
	return obj.block
}
