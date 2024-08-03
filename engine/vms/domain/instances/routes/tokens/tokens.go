package tokens

import "github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"

type tokens struct {
	hash hash.Hash
	list []Token
}

func createTokens(
	hash hash.Hash,
	list []Token,
) Tokens {
	out := tokens{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *tokens) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *tokens) List() []Token {
	return obj.list
}
