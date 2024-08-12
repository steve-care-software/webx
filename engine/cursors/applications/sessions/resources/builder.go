package resources

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type builder struct {
	builder             resources.Builder
	storageBuilder      storages.StorageBuilder
	switchersBuilder    switchers.Builder
	switcherBuilder     switchers.SwitcherBuilder
	singleBuiler        singles.Builder
	delimiterBuilder    delimiters.DelimiterBuilder
	transactionsBuilder transactions.Builder
	transactionBuilder  transactions.TransactionBuilder
	voteAdapter         signers.VoteAdapter
	dbApp               databases.Application
	hashAdapter         hash.Adapter
}

func createBuilder(
	builderIns resources.Builder,
	storageBuilder storages.StorageBuilder,
	switchersBuilder switchers.Builder,
	switcherBuilder switchers.SwitcherBuilder,
	singleBuiler singles.Builder,
	delimiterBuilder delimiters.DelimiterBuilder,
	transactionsBuilder transactions.Builder,
	transactionBuilder transactions.TransactionBuilder,
	voteAdapter signers.VoteAdapter,
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		builder:             builderIns,
		storageBuilder:      storageBuilder,
		switchersBuilder:    switchersBuilder,
		switcherBuilder:     switcherBuilder,
		singleBuiler:        singleBuiler,
		delimiterBuilder:    delimiterBuilder,
		transactionsBuilder: transactionsBuilder,
		transactionBuilder:  transactionBuilder,
		voteAdapter:         voteAdapter,
		hashAdapter:         hashAdapter,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.builder,
		app.storageBuilder,
		app.switchersBuilder,
		app.switcherBuilder,
		app.singleBuiler,
		app.delimiterBuilder,
		app.transactionsBuilder,
		app.transactionBuilder,
		app.voteAdapter,
		app.hashAdapter,
	)
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(dbApp databases.Application) Builder {
	app.dbApp = dbApp
	return app
}

// Now builds a new Application
func (app *builder) Now() (Application, error) {
	if app.dbApp == nil {
		return nil, errors.New("the database application is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.dbApp,
		app.builder,
		app.storageBuilder,
		app.switchersBuilder,
		app.switcherBuilder,
		app.singleBuiler,
		app.delimiterBuilder,
		app.transactionsBuilder,
		app.transactionBuilder,
		app.voteAdapter,
		app.hashAdapter,
	), nil
}
