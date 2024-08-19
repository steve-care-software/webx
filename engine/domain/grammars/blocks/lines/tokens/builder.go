package tokens

import (
	"errors"
	"fmt"
)

type builder struct {
	list []Token
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Token) Builder {
	app.list = list
	return app
}

// Now builds a new Tokens instance
func (app *builder) Now() (Tokens, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Tokens instance")
	}

	mp := map[string]Token{}
	for _, oneToken := range app.list {
		keyname := oneToken.Name()
		if _, ok := mp[keyname]; ok {
			str := fmt.Sprintf("the Token (name: %s) is a duplicate", keyname)
			return nil, errors.New(str)
		}
		mp[keyname] = oneToken
	}

	return createTokens(app.list, mp), nil
}
