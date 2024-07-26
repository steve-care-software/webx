package files

import (
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/databases/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/domain/entries"
)

type context struct {
	path       []string
	insertions entries.Entries
	deletions  deletes.Deletes
	pLock      *fslock.Lock
	pFile      *os.File
}
