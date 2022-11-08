package identities

import (
	"github.com/steve-care-software/webx/applications/databases/identities/selects"
	"github.com/steve-care-software/webx/domain/databases"
	"github.com/steve-care-software/webx/domain/identities"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithDtabase(database databases.Database) Builder
	Now() (Application, error)
}

// Application represents the identity application
type Application interface {
	List() ([]string, error)
	New(identity identities.Identity, password []byte) error
	Select(name string) (selects.Application, error)
}
