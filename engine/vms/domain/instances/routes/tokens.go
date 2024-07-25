package routes

import "github.com/steve-care-software/webx/engine/states/domain/hash"

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

// Remaining returns the remaining bytes
func (obj *tokens) Remaining(input []byte) []byte {
	return nil
}

// RemainingWithOmission returns the remaining (with token omission) bytes
func (obj *tokens) RemainingWithOmission(input []byte, tokenOmission Omission) []byte {
	return nil
}

// List returns the list
func (obj *tokens) List() []Token {
	return obj.list
}
