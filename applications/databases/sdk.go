package databases

import (
	"github.com/steve-care-software/webx/applications/databases/contents"
	"github.com/steve-care-software/webx/applications/databases/transactions"
)

// Application represents a database application
type Application interface {
	List() ([]string, error)
	New(name string) error
	Content(name string) (contents.Application, error)
	Transaction(name string) (transactions.Application, error)
}
