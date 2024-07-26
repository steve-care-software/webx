package applications

import (
	"github.com/steve-care-software/webx/engine/databases/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/domain/headers/containers/pointers"
)

// Application represents the database application
type Application interface {
	Begin(dbPath []string) (*uint, error)
	Retrieve(context uint, pointer pointers.Pointer) ([]byte, error)
	RetrieveAll(context uint, pointers pointers.Pointers) ([][]byte, error)
	Insert(context uint, entry entries.Entry) error
	InsertAll(context uint, entries entries.Entries) error
	Delete(context uint, pointer pointers.Pointer) error
	DeleteAll(context uint, pointers pointers.Pointers) error
	Commit(context uint) error
	Rollback(context uint) error
	Cancel(context uint) error
}
