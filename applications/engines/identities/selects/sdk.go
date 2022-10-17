package selects

import (
	"github.com/steve-care-software/syntax/applications/engines/databases"
	"github.com/steve-care-software/syntax/domain/syntax/identities"
	"github.com/steve-care-software/syntax/domain/syntax/identities/modifications"
)

// Application represents the selected identity application
type Application interface {
	Retrieve(password string) identities.Identity
	Modify(modification modifications.Modification, password string) error
	Databases(password string) (databases.Application, error)
}
