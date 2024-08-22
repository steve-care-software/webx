package instructions

import (
	"errors"
)

type tokensBuilder struct {
	list []Token
}

func createTokensBuilder() TokensBuilder {
	out := tokensBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the tokensBuilder
func (app *tokensBuilder) Create() TokensBuilder {
	return createTokensBuilder()
}

// WithList adds a list to the tokensBuilder
func (app *tokensBuilder) WithList(list []Token) TokensBuilder {
	app.list = list
	return app
}

// Now builds a new Tokens instance
func (app *tokensBuilder) Now() (Tokens, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Tokens instance")
	}

	mp := map[string][]Token{}
	for _, oneToken := range app.list {
		keyname := oneToken.Name()
		if _, ok := mp[keyname]; !ok {
			mp[keyname] = []Token{}
		}

		mp[keyname] = append(mp[keyname], oneToken)
	}

	return createTokens(app.list, mp), nil
}
