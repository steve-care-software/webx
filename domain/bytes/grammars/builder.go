package grammars

import "errors"

type builder struct {
	root     Token
	channels Channels
}

func createBuilder() Builder {
	out := builder{
		root:     nil,
		channels: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root token to the builder
func (app *builder) WithRoot(root Token) Builder {
	app.root = root
	return app
}

// WithChannels add channels token to the builder
func (app *builder) WithChannels(channels Channels) Builder {
	app.channels = channels
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.root == nil {
		return nil, errors.New("the root Token is mandatory in order to build a Grammar instance")
	}

	if app.channels != nil {
		return createGrammarWithChannels(app.root, app.channels), nil
	}

	return createGrammar(app.root), nil
}
