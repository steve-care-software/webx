package grammars

import "github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"

type block struct {
	hash  hash.Hash
	lines []Line
}

func createBlock(
	hash hash.Hash,
	lines []Line,
) Block {
	out := block{
		hash:  hash,
		lines: lines,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Lines returns the lines
func (obj *block) Lines() []Line {
	return obj.lines
}
