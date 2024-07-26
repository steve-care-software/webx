package applications

import (
	"github.com/steve-care-software/webx/engine/databases/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/domain/headers/states/containers/pointers"
	"github.com/steve-care-software/webx/engine/databases/domain/retrievals"
)

// Application represents the database application
type Application interface {
	Begin(path []string) (*uint, error)
	List(context uint, keyname string, index uint, length uint) (pointers.Pointer, error)
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
