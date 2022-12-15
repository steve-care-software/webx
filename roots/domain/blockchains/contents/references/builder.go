package references

type builder struct {
	contentFactory ContentFactory
	content        Content
	blockchain     Blockchain
}

func createBuilder(
	contentFactory ContentFactory,
) Builder {
	out := builder{
		contentFactory: contentFactory,
		content:        nil,
		blockchain:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.contentFactory,
	)
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// WithBlockchain adds a blockchain to the builder
func (app *builder) WithBlockchain(blockchain Blockchain) Builder {
	app.blockchain = blockchain
	return app
}

// Now builds a new Reference instance
func (app *builder) Now() (Reference, error) {
	if app.content == nil {
		content, err := app.contentFactory.Create()
		if err != nil {
			return nil, err
		}

		app.content = content
	}

	if app.blockchain != nil {
		return createReferenceWithBlockchain(app.content, app.blockchain), nil
	}

	return createReference(app.content), nil
}
