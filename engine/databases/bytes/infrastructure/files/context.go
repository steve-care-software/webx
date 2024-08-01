package files

import (
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
)

type context struct {
	path          []string
	name          string
	pDataIndex    *uint64
	currentHeader states.States
	insertions    entries.Entries
	deletions     delimiters.Delimiters
	pFile         *os.File
	pLock         *fslock.Lock
}
