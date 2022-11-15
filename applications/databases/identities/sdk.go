package identities

import (
	"github.com/steve-care-software/webx/applications/databases/identities/selects"
	"github.com/steve-care-software/webx/domain/identities"
)

// Application represents the identity application
type Application interface {
	List() ([]string, error)
	New(identity identities.Identity, password []byte) error
	Select(name string) (selects.Application, error)
}
