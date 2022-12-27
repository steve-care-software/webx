package grammars

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

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

// Points returns the amount of points a block contains
func (obj *block) Points() uint {
	amount := uint(0)
	for _, oneLine := range obj.lines {
		amount += oneLine.Points()
	}

	return amount
}

// Lines returns the lines
func (obj *block) Lines() []Line {
	return obj.lines
}
