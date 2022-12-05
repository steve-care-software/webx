package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type tokensBuilder struct {
	hashAdapter hash.Adapter
	list        []Token
}

func createTokensBuilder(
	hashAdapter hash.Adapter,
) TokensBuilder {
	out := tokensBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokensBuilder) Create() TokensBuilder {
	return createTokensBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
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

	data := [][]byte{}
	for _, oneToken := range app.list {
		data = append(data, oneToken.Hash().Bytes())
	}

	phash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createTokens(*phash, app.list), nil
}
