package instructions

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

type instruction struct {
	hash    hash.Hash
	content Content
}

func createInstruction(
	hash hash.Hash,
	content Content,
) Instruction {
	out := instruction{
		hash:    hash,
		content: content,
	}

	return &out
}

// Hash returns the hash
func (obj *instruction) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *instruction) Content() Content {
	return obj.content
}
