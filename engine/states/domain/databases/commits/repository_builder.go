package commits

type repositoryBuilder struct {
	innerPath []string
	basePath  []string
}

func createRepositoryBuilder(
	innerPath []string,
) RepositoryBuilder {
	out := repositoryBuilder{
		innerPath: innerPath,
		basePath:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *repositoryBuilder) Create() RepositoryBuilder {
	return createRepositoryBuilder(
		app.innerPath,
	)
}

// WithBasePath adds a base path to the builder
func (app *repositoryBuilder) WithBasePath(basePath []string) RepositoryBuilder {
	app.basePath = basePath
	return app
}

// Now builds a new Repository instance
func (app *repositoryBuilder) Now() (Repository, error) {
	return nil, nil
}
