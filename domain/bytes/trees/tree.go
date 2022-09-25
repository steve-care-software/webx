package trees

import "github.com/steve-care-software/syntax/domain/bytes/grammars"

type tree struct {
	grammar grammars.Token
	block   Block
	suffix  Trees
}

func createTree(
	grammar grammars.Token,
	block Block,
) Tree {
	return createTreeInternally(grammar, block, nil)
}

func createTreeWithSuffix(
	grammar grammars.Token,
	block Block,
	suffix Trees,
) Tree {
	return createTreeInternally(grammar, block, suffix)
}

func createTreeInternally(
	grammar grammars.Token,
	block Block,
	suffix Trees,
) Tree {
	out := tree{
		grammar: grammar,
		block:   block,
		suffix:  suffix,
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

// HasSuffix returns true if there is suffix, false otherwise
func (obj *tree) HasSuffix() bool {
	return obj.suffix != nil
}

// Suffix returns the block
func (obj *tree) Suffix() Trees {
	return obj.suffix
}
