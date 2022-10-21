package identities

import (
	"github.com/steve-care-software/webx/applications/identities/selects"
	"github.com/steve-care-software/webx/domain/identities"
)

type application struct {
	selectAppBuilder selects.Builder
	repository       identities.Repository
	service          identities.Service
}

func createApplication(
	selectAppBuilder selects.Builder,
	repository identities.Repository,
	service identities.Service,
) Application {
	out := application{
		selectAppBuilder: selectAppBuilder,
		repository:       repository,
		service:          service,
	}

	return &out
}

// List lists the identity names
func (app *application) List() ([]string, error) {
	return app.repository.List()
}

// New saves a new identity
func (app *application) New(identity identities.Identity, password string) error {
	return app.service.Insert(identity, password)
}

// Select returns the select identity application
func (app *application) Select(name string) (selects.Application, error) {
	return app.selectAppBuilder.Create().
		WithName(name).
		Now()
}
