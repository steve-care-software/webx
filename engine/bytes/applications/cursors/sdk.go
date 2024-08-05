package cursors

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/cursors"
	"github.com/steve-care-software/webx/engine/bytes/domain/cursors/status"
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/originals"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCursor(cursor cursors.Cursor) Builder
	Now() (Application, error)
}

// Application represents the cursor application
type Application interface {
	Cursor() cursors.Cursor            // returns the current cursor
	Records() (status.Statuses, error) // returns the recorded cursor status representations
	Erase(name string) error           // erase a cursor using its name
	Record(name string) error          // record the cursor to this name
	Replace(name string) error         // put the cursor to what the name was pointing to
	Set(name string) error             // set the cursor to this element (horizontally)
	Down(name string) error            // set the cursor to this element (sub-element or 'down')
	Climb(name string) error           // set the cursor to this element (parent-element or 'climb')
	Insert(original originals.Original) error
	Update(original string, updated originals.Original) error
	Delete(name string) error
	Recover(name string) error
	Purge(name string) error
	PurgeAll() error
	Move(name string, devName string, deleteOriginal bool) error // moves a development iteration to a production iteration inside the current iteration
	MetaData() (delimiters.Delimiter, error)                     // returns the current branch meta data
	Merge(deleteOriginal bool) error                             // merge the current branch into its parent and set its parent as the current branch

	// data:
	InsertData(data []byte) (delimiters.Delimiter, error)
	UpdateData(original delimiters.Delimiter, updated []byte) error
	DeleteData(delete delimiters.Delimiter) error
}
