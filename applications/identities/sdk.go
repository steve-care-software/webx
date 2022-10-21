package identities

import (
	"github.com/steve-care-software/webx/applications/identities/selects"
	"github.com/steve-care-software/webx/domain/identities"
)

// NewApplication creates a new application
func NewApplication(
	repository identities.Repository,
	service identities.Service,
) Application {
	selectAppBuilder := selects.NewBuilder(repository, service)
	return createApplication(selectAppBuilder, repository, service)
}

// Application represents an identity application
type Application interface {
	List() ([]string, error)
	New(identity identities.Identity, password string) error
	Select(name string) (selects.Application, error)
}
