package references

type builder struct {
	contentFactory ContentFactory
	content        Content
	commits        Commits
}

func createBuilder(
	contentFactory ContentFactory,
) Builder {
	out := builder{
		contentFactory: contentFactory,
		content:        nil,
		commits:        nil,
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

// WithCommits add commits to the builder
func (app *builder) WithCommits(commits Commits) Builder {
	app.commits = commits
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

	if app.commits != nil {
		return createReferenceWithCommits(app.content, app.commits), nil
	}

	return createReference(app.content), nil
}
