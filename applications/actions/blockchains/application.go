package blockchains

import (
	"github.com/steve-care-software/syntax/applications/actions/blockchains/selects"
	"github.com/steve-care-software/syntax/domain/blockchains"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

type application struct {
	builder    selects.Builder
	repository blockchains.Repository
}

func createApplication(
	builder selects.Builder,
	repository blockchains.Repository,
) Application {
	out := application{
		builder:    builder,
		repository: repository,
	}

	return &out
}

// List lists the blockchain references
func (app *application) List() ([]hash.Hash, error) {
	return app.repository.List()
}

// Select selects a blockchain application by reference
func (app *application) Select(ref hash.Hash) (selects.Application, error) {
	return app.builder.Create().WithReference(ref).Now()
}
