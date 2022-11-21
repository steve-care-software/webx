package databases

import (
	"github.com/steve-care-software/webx/applications/databases"
	"github.com/steve-care-software/webx/applications/databases/contents"
	"github.com/steve-care-software/webx/applications/databases/transactions"
)

type application struct {
	contentAppBuilder contents.Builder
	trxAppBuilder     transactions.Builder
	dirPath           string
}

func createApplication(
	contentAppBuilder contents.Builder,
	trxAppBuilder transactions.Builder,
	dirPath string,
) databases.Application {
	out := application{
		contentAppBuilder: contentAppBuilder,
		trxAppBuilder:     trxAppBuilder,
		dirPath:           dirPath,
	}

	return &out
}

// List lists the databases
func (app *application) List() ([]string, error) {
	return nil, nil
}

// New creates a new database by name
func (app *application) New(name string) error {
	return nil
}

// Content returns a content application using the given name
func (app *application) Content(name string) (contents.Application, error) {
	return app.contentAppBuilder.Create().
		WithName(name).
		Now()
}

// Transaction returns a transaction application using the given name
func (app *application) Transaction(name string) (transactions.Application, error) {
	return app.trxAppBuilder.Create().
		WithName(name).
		Now()
}
