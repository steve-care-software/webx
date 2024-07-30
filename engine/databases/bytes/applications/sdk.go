package applications

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/listers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithBasePath(basePath []string) Builder
	Now() (Application, error)
}

// Application represents the database application
type Application interface {
	Begin(name string) (*uint, error)
	List(context uint, lister listers.Lister) (retrievals.Retrievals, error)
	Amount(context uint, keyname string) (*uint, error)
	Retrieve(context uint, retrival retrievals.Retrieval) ([]byte, error)
	RetrieveAll(context uint, retrievals retrievals.Retrievals) ([][]byte, error)
	Insert(context uint, entry entries.Entry) error
	InsertAll(context uint, entries entries.Entries) error
	Delete(context uint, delete deletes.Delete) error
	DeleteAll(context uint, deletes deletes.Deletes) error
	Commit(context uint) error
	Rollback(context uint) error
	RollbackTo(context uint, amount uint) error
	RollFront(context uint) error
	RollFrontTo(context uint, amount uint) error
	States(context uint, includesDeleted bool) (*uint, error)
	DeletedStates(context uint) (*uint, error)
	Cancel(context uint) error
	Purge(context uint) error
}
