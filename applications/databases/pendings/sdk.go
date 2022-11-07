package pendings

import (
	"github.com/steve-care-software/webx/domain/databases"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Application represents the pendings application
type Application interface {
	List() (entities.Identifiers, error)
	Retrieve(identifier entities.Identifier) (databases.Entry, error)
}
