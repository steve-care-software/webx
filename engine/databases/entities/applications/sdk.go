package applications

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/entities"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithBasePath(basePath []string) Builder
	WithNamespace(namespace string) Builder
	Now() (Application, error)
}

// Application represents the database application
type Application interface {
	Begin(name string) (*uint, error)
	//List(context uint, lister listers.Lister) ([]hash.Hash, error)
	//Amount(context uint, keyname string) (*uint, error)
	Retrieve(context uint, hash hash.Hash) (entities.Entity, error)
	RetrieveAll(context uint, hashes []hash.Hash) ([]entities.Entity, error)
	Insert(context uint, entity entities.Entity) error
	InsertAll(context uint, entities []entities.Entity) error
	Delete(context uint, hash hash.Hash) error
	DeleteAll(context uint, hashes []hash.Hash) error
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
