package instructions

import (
	"errors"
	"fmt"
)

type tokens struct {
	list []Token
	mp   map[string][]Token
}

func createTokens(
	list []Token,
	mp map[string][]Token,
) Tokens {
	out := tokens{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of token
func (obj *tokens) List() []Token {
	return obj.list
}

// Fetch fetches a token by name and index
func (obj *tokens) Fetch(name string, idx uint) (Token, error) {
	if ins, ok := obj.mp[name]; ok {
		length := uint(len(ins))
		if idx >= length {
			str := fmt.Sprintf("the token (%s) could not be found at index (%d), its length is: %d", name, idx, length)
			return nil, errors.New(str)
		}

		return ins[idx], nil
	}

	str := fmt.Sprintf("the token (name: %s) does not exists", name)
	return nil, errors.New(str)
}
