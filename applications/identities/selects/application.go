package selects

import (
	"github.com/steve-care-software/webx/domain/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
)

type application struct {
	builder              identities.Builder
	modificationsBuilder modifications.Builder
	repository           identities.Repository
	service              identities.Service
	name                 string
}

func createApplication(
	builder identities.Builder,
	modificationsBuilder modifications.Builder,
	repository identities.Repository,
	service identities.Service,
	name string,
) Application {
	out := application{
		builder:              builder,
		modificationsBuilder: modificationsBuilder,
		repository:           repository,
		service:              service,
		name:                 name,
	}

	return &out
}

// Retrieve retrieves an identity
func (app *application) Retrieve(password string) (identities.Identity, error) {
	return app.repository.Retrieve(app.name, password)
}

// Modify modifies an identity
func (app *application) Modify(modification modifications.Modification, currentPassword string, newPassword string) error {
	identity, err := app.Retrieve(currentPassword)
	if err != nil {
		return err
	}

	modificationsList := identity.Modifications().List()
	modificationsList = append(modificationsList, modification)
	modifications, err := app.modificationsBuilder.Create().WithList(modificationsList).Now()
	if err != nil {
		return err
	}

	updatedIns, err := app.builder.Create().
		WithModifications(modifications).
		Now()

	if err != nil {
		return err
	}

	return app.service.Update(updatedIns, currentPassword, newPassword)
}
