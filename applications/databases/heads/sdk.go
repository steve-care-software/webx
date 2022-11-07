package heads

import "github.com/steve-care-software/webx/domain/databases"

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithDtabase(database databases.Database) Builder
	Now() (Application, error)
}

// Application represents the head application
type Application interface {
	Retrieve() databases.Head
	Migrate(migration databases.Migration) error
}
