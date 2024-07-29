package files

import (
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
)

type context struct {
	path          []string
	currentHeader states.States
	insertions    entries.Entries
	deletions     deletes.Deletes
	pLock         *fslock.Lock
	pFile         *os.File
}
