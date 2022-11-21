package databases

import (
	"errors"

	"github.com/steve-care-software/webx/applications/databases"
	"github.com/steve-care-software/webx/applications/databases/contents"
	"github.com/steve-care-software/webx/applications/databases/transactions"
)

type builder struct {
	contentAppBuilder contents.Builder
	trxAppBuilder     transactions.Builder
	dirPath           string
}

func createBuilder(
	contentAppBuilder contents.Builder,
	trxAppBuilder transactions.Builder,
) Builder {
	out := builder{
		contentAppBuilder: contentAppBuilder,
		trxAppBuilder:     trxAppBuilder,
		dirPath:           "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.contentAppBuilder,
		app.trxAppBuilder,
	)
}

// WithDirPath adds a dirPath to the builder
func (app *builder) WithDirPath(dirPath string) Builder {
	app.dirPath = dirPath
	return app
}

// Now builds a new Database Application
func (app *builder) Now() (databases.Application, error) {
	if app.dirPath == "" {
		return nil, errors.New("the dirPath is mandatory in order to build a database application")
	}

	return createApplication(
		app.contentAppBuilder,
		app.trxAppBuilder,
		app.dirPath,
	), nil
}
