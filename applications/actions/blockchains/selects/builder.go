package selects

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/blockchains"
	"github.com/steve-care-software/syntax/domain/blockchains/blocks"
	"github.com/steve-care-software/syntax/domain/blockchains/transactions"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

type builder struct {
	builder             blockchains.Builder
	repository          blockchains.Repository
	service             blockchains.Service
	blockRepository     blocks.Repository
	transactionsBuilder transactions.Builder
	pReference          *hash.Hash
}

func createBuilder(
	blockchainBuilder blockchains.Builder,
	repository blockchains.Repository,
	service blockchains.Service,
	blockRepository blocks.Repository,
	transactionsBuilder transactions.Builder,
) Builder {
	out := builder{
		builder:             blockchainBuilder,
		repository:          repository,
		service:             service,
		blockRepository:     blockRepository,
		transactionsBuilder: transactionsBuilder,
		pReference:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.builder,
		app.repository,
		app.service,
		app.blockRepository,
		app.transactionsBuilder,
	)
}

// WithReference adds a reference to the builder
func (app *builder) WithReference(ref hash.Hash) Builder {
	app.pReference = &ref
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.pReference == nil {
		return nil, errors.New("the reference hash is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.builder,
		app.repository,
		app.service,
		app.blockRepository,
		app.transactionsBuilder,
		*app.pReference,
	), nil
}
