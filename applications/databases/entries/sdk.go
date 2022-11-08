package entries

import (
	"github.com/steve-care-software/webx/domain/databases"
)

// Application represents the pendings application
type Application interface {
	Retrieve(pointer databases.Pointer) (databases.Entry, error)
	Insert(entry databases.Entry) (databases.Pointer, error)
}
