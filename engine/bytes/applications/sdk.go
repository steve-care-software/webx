package applications

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/cursors/status"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
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
	Status(context uint) (status.Status, error)    // returns the status, which is the cursor without the lists
	Records(context uint) (status.Statuses, error) // returns the recorded cursor status representations
	Erase(context uint, cursorName string) error   // erase a cursor using its name
	Record(context uint, name string) error        // record the cursor to this name
	Replace(context uint, cursorName string) error // put the cursor to what the name was pointing to
	Retrieve(context uint, name string, includesDeleted bool) (originals.Original, error)
	RetrieveAll(context uint, includesDeleted bool) (originals.Originals, error)
	Deleted(context uint) (originals.Originals, error)
	Set(context uint, name string) error   // set the cursor to this element (horizontally)
	Down(context uint, name string) error  // set the cursor to this element (sub-element or 'down')
	Climb(context uint, name string) error // set the cursor to this element (parent-element or 'climb')
	Insert(context uint, original originals.Original) error
	Update(context uint, original string, updated originals.Original) error
	Delete(context uint, name string) error
	Recover(context uint, name string) error
	Purge(context uint, name string) error
	PurgeAll(context uint) error
	Cleanup(context uint) error
	Move(context uint, name string, devName string, deleteOriginal bool) error // moves a development iteration to a production iteration inside the current iteration
	MetaData(context uint) (delimiters.Delimiter, error)                       // returns the current branch meta data
	Merge(context uint, deleteOriginal bool) error                             // merge the current branch into its parent and set its parent as the current branch
	Commit(context uint) error
	CommitWithMetaData(context uint, metaData []byte) error
	Close(context uint) error

	// data:
	RetrieveData(context uint, retrival delimiters.Delimiter) ([]byte, error)
	InsertData(context uint, data []byte) (delimiters.Delimiter, error)
	UpdateData(context uint, original delimiters.Delimiter, updated []byte) error
	DeleteData(context uint, delete delimiters.Delimiter) error
}
