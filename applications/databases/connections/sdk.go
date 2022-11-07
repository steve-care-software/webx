package connections

import (
	"net/url"

	"github.com/steve-care-software/webx/domain/databases"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithDtabase(database databases.Database) Builder
	Now() (Application, error)
}

// Application represents the connection application
type Application interface {
	List() []url.URL
	Connect(url url.URL) error
	ConnectList(urls []url.URL) error
}
