package tokens

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

type token struct {
	hash  hash.Hash
	lines Lines
}

func createToken(
	hash hash.Hash,
	lines Lines,
) Token {
	return createTokenInternally(hash, lines)
}

func createTokenInternally(
	hash hash.Hash,
	lines Lines,
) Token {
	out := token{
		hash:  hash,
		lines: lines,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// Lines returns the lines
func (obj *token) Lines() Lines {
	return obj.lines
}
