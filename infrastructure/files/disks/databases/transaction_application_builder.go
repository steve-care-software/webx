package databases

import (
	"errors"

	"github.com/steve-care-software/webx/applications/databases/transactions"
)

type transactionApplicationBuilder struct {
	name string
}

func createTransactionApplicationBuilder() transactions.Builder {
	out := transactionApplicationBuilder{}
	return &out
}

// Create initializes the builder
func (app *transactionApplicationBuilder) Create() transactions.Builder {
	return createTransactionApplicationBuilder()
}

// WithName adds a name to the builder
func (app *transactionApplicationBuilder) WithName(name string) transactions.Builder {
	app.name = name
	return app
}

// Now builds a new transaction application
func (app *transactionApplicationBuilder) Now() (transactions.Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a transaction application")
	}

	return createTransactionApplication(app.name), nil
}
