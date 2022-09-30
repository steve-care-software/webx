package grammars

type tokens struct {
	list []Token
}

func createTokens(
	list []Token,
) Tokens {
	out := tokens{
		list: list,
	}

	return &out
}

// List returns the list of tokens
func (obj *tokens) List() []Token {
	return obj.list
}
